package main

import (
	"encoding/json"
	"gorilla/mux"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"

	types "arqui-reports/types"
)

var sampleUserReport = types.UserReport{
	GoodHabits: []types.Habit{
		{
			ID:         0,
			UserID:     0,
			Name:       "Good habit name",
			HabitType:  0,
			Difficulty: 0,
			Score:      0,
			Color:      0,
		},
	},
	BadHabits: []types.Habit{
		{
			ID:         0,
			UserID:     0,
			Name:       "Bad habit name",
			HabitType:  0,
			Difficulty: 0,
			Score:      0,
			Color:      0,
		},
	},
	TodayTasks: []types.Task{
		{
			ID:             0,
			UserID:         0,
			Title:          "Today task name",
			Description:    "Today task description",
			Difficulty:     "Today task difficulty",
			DueDate:        "Today task due date",
			CompletionDate: "Completion",
			Reminder:       "Today task reminder",
			Done:           true,
		},
	},
	DelayedTasks: []types.Task{
		{
			ID:             0,
			UserID:         0,
			Title:          "Delayed task name",
			Description:    "Delayed task description",
			Difficulty:     "Delayed task difficulty",
			DueDate:        "Delayed task due date",
			CompletionDate: "Completion",
			Reminder:       "Delayed task reminder",
			Done:           true,
		},
	},
}

func getJSON(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func userReportHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	if userID == "" {
		w.WriteHeader(http.StatusForbidden)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	} else {

		tasks := []types.Task{}
		getJSON("https://habittonapigateway.herokuapp.com/users/tasks?userId="+userID, &tasks)
		habits := []types.Habit{}
		getJSON("https://habittonapigateway.herokuapp.com/users/habits?userId="+userID, &habits)

		goodHabits := []types.Habit{}
		badHabits := []types.Habit{}
		tasksToday := []types.Task{}
		tasksDelayed := []types.Task{}

		for _, habit := range habits {
			if habit.Score >= 50 {
				goodHabits = append(goodHabits, habit)
			} else if habit.Score < 0 {
				badHabits = append(badHabits, habit)
			}
		}

		layout := "02/01/2006"
		todayDateString := time.Now().Format(layout)
		todayDate, err := time.Parse(layout, todayDateString)
		if err != nil {
			log.Panicln(err)
		}
		for _, task := range tasks {
			if !task.Done {
				dueDateString := task.DueDate
				dueDate, err := time.Parse(layout, dueDateString)
				if err != nil {
					log.Panicln(err)
				}

				if dueDate.Equal(todayDate) {
					tasksToday = append(tasksToday, task)
				}
				if dueDate.Before(todayDate) {
					tasksDelayed = append(tasksDelayed, task)
				}
			}
		}

		var response = types.UserReport{
			BadHabits:    badHabits,
			GoodHabits:   goodHabits,
			TodayTasks:   tasksToday,
			DelayedTasks: tasksDelayed,
		}

		userReportJSON, err := json.Marshal(response)
		if err != nil {
			log.Panicln(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(userReportJSON)
	}
}

func adminTasksReport(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	if userID != "0" {
		w.WriteHeader(http.StatusForbidden)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	} else {

		tasks := []types.Task{}
		getJSON("https://habittonapigateway.herokuapp.com/admin/tasks/?userId=0", &tasks)

		completedTotal := 0
		completedBefore := 0
		completedAfter := 0
		delayed := 0
		availableTotal := 0
		availableRemainig := 0
		availableForToday := 0

		layout := "02/01/2006"
		todayDateString := time.Now().Format(layout)
		todayDate, err := time.Parse(layout, todayDateString)
		if err != nil {
			log.Panicln(err)
		}
		for _, task := range tasks {

			dueDateString := task.DueDate
			dueDate, err := time.Parse(layout, dueDateString)
			if err != nil {
				log.Panicln(err)
			}

			if task.Done {
				completionDateString := task.CompletionDate
				completionDate, err := time.Parse(layout, completionDateString)
				if err != nil {
					log.Panicln(err)
				}

				if completionDate.After(dueDate) || completionDate.Equal(dueDate) {
					completedBefore++
				} else {
					completedAfter++
				}
				completedTotal++
			} else if dueDate.Before(todayDate) {
				delayed++
			} else {
				if dueDate.Equal(todayDate) {
					availableForToday++
				} else {
					availableRemainig++
				}
				availableTotal++
			}
		}

		var response = types.AdminTasksReport{
			Completed: types.CompletedTasks{
				Total:  completedTotal,
				Before: completedBefore,
				After:  completedAfter,
			},
			Delayed: delayed,
			Available: types.AvailableTasks{
				Total:     availableTotal,
				Remaining: availableRemainig,
				ForToday:  availableForToday,
			},
		}
		adminTasksReportJSON, err := json.Marshal(response)
		if err != nil {
			log.Panicln(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(adminTasksReportJSON)
	}

}

func adminHabitsReport(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	if userID != "0" {
		w.WriteHeader(http.StatusForbidden)
	} else {

		habits := []types.Habit{}
		getJSON("https://arquitectura-habits.herokuapp.com/admin/habits?userId=0", &habits)

		countRed := 0
		countOrange := 0
		countYellow := 0
		countGreen := 0
		countBlue := 0
		highestScoreUserID := -1
		lowestScoreUserID := -1
		highestScoreName := ""
		lowestScoreName := ""
		highestScore := math.Inf(-1)
		lowestScore := math.Inf(1)

		for _, habit := range habits {
			if float64(habit.Score) < lowestScore {
				lowestScore = float64(habit.Score)
				lowestScoreUserID = habit.UserID
				lowestScoreName = habit.Name
			}
			if float64(habit.Score) > highestScore {
				highestScore = float64(habit.Score)
				highestScoreUserID = habit.UserID
				highestScoreName = habit.Name
			}

			if habit.Score < 0 {
				countRed++
			} else if habit.Score >= 0 && habit.Score < 10 {
				countOrange++
			} else if habit.Score >= 10 && habit.Score < 40 {
				countYellow++
			} else if habit.Score >= 40 && habit.Score < 50 {
				countGreen++
			} else if habit.Score >= 50 {
				countBlue++
			}
		}

		worstHabit := types.HabitOwner{}
		if lowestScoreUserID != -1 {
			userData := types.UserData{}
			getJSON("https://habittonapigateway.herokuapp.com/admin/users/name?userId=0&searchUserId="+strconv.Itoa(lowestScoreUserID), &userData)
			worstHabit.Name = lowestScoreName
			worstHabit.Username = userData.Name
		} else {
			worstHabit.Name = ""
			worstHabit.Username = ""
		}

		bestHabit := types.HabitOwner{}
		if lowestScoreUserID != -1 {
			userData := types.UserData{}
			getJSON("https://habittonapigateway.herokuapp.com/admin/users/name?userId=0&searchUserId="+strconv.Itoa(highestScoreUserID), &userData)
			bestHabit.Name = highestScoreName
			bestHabit.Username = userData.Name
		} else {
			bestHabit.Name = ""
			bestHabit.Username = ""
		}

		var response = types.AdminHabitsReport{
			BestHabit:  bestHabit,
			WorstHabit: worstHabit,
			PerRange: types.Ranges{
				Red:    countRed,
				Orange: countOrange,
				Yellow: countYellow,
				Green:  countGreen,
				Blue:   countBlue,
			},
		}

		adminHabitsReportJSON, err := json.Marshal(response)
		if err != nil {
			log.Panicln(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(adminHabitsReportJSON)
	}
}

func aliveHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func determineListenPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return ":5000", nil
	}
	return ":" + port, nil
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users/reports", userReportHandler)
	r.HandleFunc("/admin/reports/tasks", adminTasksReport)
	r.HandleFunc("/admin/reports/habits", adminHabitsReport)
	r.HandleFunc("/", aliveHandler)

	port, err := determineListenPort()
	if err != nil {
		log.Fatalln(err)
	}
	http.ListenAndServe(port, r)
}
