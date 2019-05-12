package main

import (
	"encoding/json"
	"gorilla/mux"
	"log"
	"net/http"
	"os"
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
		},
	},
	TodayTasks: []types.Task{
		{
			ID:          0,
			UserID:      0,
			Title:       "Today task name",
			Description: "Today task description",
			Difficulty:  "Today task difficulty",
			DueDate:     "Today task due date",
			Reminder:    "Today task reminder",
			Done:        true,
		},
	},
	DelayedTasks: []types.Task{
		{
			ID:          0,
			UserID:      0,
			Title:       "Delayed task name",
			Description: "Delayed task description",
			Difficulty:  "Delayed task difficulty",
			DueDate:     "Delayed task due date",
			Reminder:    "Delayed task reminder",
			Done:        true,
		},
	},
}

var sampleAdminHabitsReport = types.AdminHabitsReport{
	PerRange: types.Ranges{
		Red:    0,
		Orange: 0,
		Yellow: 0,
		Green:  0,
		Blue:   0,
	},
	WorstHabit: types.HabitOwner{
		Name:     "Name Worst",
		Username: "username_worst",
	},
	BestHabit: types.HabitOwner{
		Name:     "Name Best",
		Username: "username_best",
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
		userReportJSON, err := json.Marshal(sampleUserReport)
		if err != nil {
			log.Println(err)
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
		log.Println(tasks)

		completedTotal := 0
		completedBefore := 0
		completedAfter := 0
		delayed := 0
		availableTotal := 0
		availableRemainig := 0
		availableForToday := 0

		layout := "02/01/2006"
		for _, task := range tasks {
			todayDateString := time.Now().Format(layout)
			todayDate, err := time.Parse(layout, todayDateString)
			if err != nil {
				log.Println(err)
			}
			dueDateString := task.DueDate
			dueDate, err := time.Parse(layout, dueDateString)
			if err != nil {
				log.Println(err)
			}

			if task.Done {
				completionDateString := task.CompletionDate
				completionDate, err := time.Parse(layout, completionDateString)
				if err != nil {
					log.Println(err)
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
			log.Println(err)
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
		adminHabitsReportJSON, err := json.Marshal(sampleAdminHabitsReport)
		if err != nil {
			log.Println(err)
		}

		habits := []types.Habit{}
		getJSON("https://arquitectura-habits.herokuapp.com/admin/habits?userId=0", &habits)
		log.Println(habits)

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
