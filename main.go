package main

import (
	"arqui-reports/handlers"
	"gorilla/mux"
	"log"
	"net/http"
	"os"
)

func determineListenPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return ":5000", nil
	}
	return ":" + port, nil
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users/reports", handlers.UserReportHandler)
	r.HandleFunc("/admin/reports/tasks", handlers.AdminTasksReport)
	r.HandleFunc("/admin/reports/habits", handlers.AdminHabitsReport)
	r.HandleFunc("/", handlers.AliveHandler)

	port, err := determineListenPort()
	if err != nil {
		log.Fatalln(err)
	}
	http.ListenAndServe(port, r)
}
