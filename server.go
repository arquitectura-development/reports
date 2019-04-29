package reports

import (
	"gorilla/mux"
	"net/http"
)

func userReportHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
}

func adminTasksReport(w http.ResponseWriter, r *http.Request) {

}

func adminHabitsReport(w http.ResponseWriter, r *http.Request) {
}

func main() {
	router := mux.NewRouter()
	router.Methods("GET")
	router.HandleFunc("/reports/users/{userId}", userReportHandler)
	router.HandleFunc("/reports/admin/tasks", adminTasksReport)
	router.HandleFunc("/reports/admin/habits", adminHabitsReport)
	http.Handle("/", router)
}
