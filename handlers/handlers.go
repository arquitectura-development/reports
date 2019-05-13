package handlers

import (
	"arqui-reports/types"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func getJSON(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func UserReportHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	if userID == "" {
		w.WriteHeader(http.StatusForbidden)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	} else {

		tasks := []types.Task{}
		getJSON("https://habittonapigateway.herokuapp.com/users/tasks?userId="+userID, &tasks)
		habits := []types.Habit{}
		getJSON("https://habittonapigateway.herokuapp.com/users/habits?userId="+userID, &habits)

		response := types.UserReport{}

		ProcessUserReport(tasks, habits, &response)

		userReportJSON, err := json.Marshal(response)
		if err != nil {
			log.Panicln(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(userReportJSON)
	}
}

func AdminTasksReport(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	if userID != "0" {
		w.WriteHeader(http.StatusForbidden)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	} else {

		tasks := []types.Task{}
		getJSON("https://habittonapigateway.herokuapp.com/admin/tasks/?userId=0", &tasks)

		response := types.AdminTasksReport{}
		ProcessAdminTasks(tasks, &response)

		adminTasksReportJSON, err := json.Marshal(response)
		if err != nil {
			log.Panicln(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(adminTasksReportJSON)
	}

}

func AdminHabitsReport(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	if userID != "0" {
		w.WriteHeader(http.StatusForbidden)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	} else {

		habits := []types.Habit{}
		getJSON("https://arquitectura-habits.herokuapp.com/admin/habits?userId=0", &habits)

		lowestTemp := types.UserHabitTemp{UserID: -1}
		highestTemp := types.UserHabitTemp{UserID: -1}
		ranges := types.Ranges{}

		ProcessAdminHabitsWithoutNames(habits, &ranges, &lowestTemp, &highestTemp)

		worstHabit := types.HabitOwner{}
		if lowestTemp.UserID != -1 {
			userData := types.UserData{}
			getJSON("https://habittonapigateway.herokuapp.com/admin/users/name?userId=0&searchUserId="+strconv.Itoa(lowestTemp.UserID), &userData)
			worstHabit.Name = lowestTemp.HabitName
			worstHabit.Username = userData.Name
		} else {
			worstHabit.Name = ""
			worstHabit.Username = ""
		}

		bestHabit := types.HabitOwner{}
		if highestTemp.UserID != -1 {
			userData := types.UserData{}
			getJSON("https://habittonapigateway.herokuapp.com/admin/users/name?userId=0&searchUserId="+strconv.Itoa(highestTemp.UserID), &userData)
			bestHabit.Name = highestTemp.HabitName
			bestHabit.Username = userData.Name
		} else {
			bestHabit.Name = ""
			bestHabit.Username = ""
		}

		var response = types.AdminHabitsReport{
			BestHabit:  bestHabit,
			WorstHabit: worstHabit,
			PerRange:   ranges,
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

func AliveHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
