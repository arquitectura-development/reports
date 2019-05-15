package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var userID = "userId"

func TestAliveHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AliveHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("AliveHandler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestAdminHabitsReportValid(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin/reports/habits", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add(userID, "0")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AdminHabitsReportHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("AdminHabitsReport returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestAdminHabitsReportInvalid(t *testing.T) {

	req, err := http.NewRequest("GET", "/admin/reports/habits", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add(userID, "1")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AdminHabitsReportHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("AdminHabitsReport returned wrong status code: got %v want %v", status, http.StatusForbidden)
	}
}

func TestAdminTasksReportValid(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin/reports/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add(userID, "0")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AdminTasksReportHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("AdminTasksReport returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestAdminTasksReportInvalid(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin/reports/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add(userID, "1")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AdminTasksReportHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("AdminTasksReport returned wrong status code: got %v want %v", status, http.StatusForbidden)
	}
}

func TestUserReportValid(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/reports", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add(userID, "1")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UserReportHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("UserReportHandler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestUserHandlerInvalid(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/reports", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add(userID, "")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UserReportHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("UserReportHandler returned wrong status code: got %v want %v", status, http.StatusForbidden)
	}
}
