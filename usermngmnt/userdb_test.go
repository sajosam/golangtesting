package usermngmnt

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"usermngmnt"
)


func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(usermngmnt.HealthCheck)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("HealthCheck handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := "Super Secret Area"
	if rr.Body.String() != expected {
		t.Errorf("HealthCheck handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(usermngmnt.GetUser)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetUser handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestAddUser(t *testing.T) {
	req, err := http.NewRequest("POST", "/adduser", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(usermngmnt.AddUser)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("AddUser handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetUserInd(t *testing.T) {
	req, err := http.NewRequest("GET", "/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(usermngmnt.GetUserInd)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetUserInd handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDelUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/delUser/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(usermngmnt.DelUser)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("DelUser handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUpdateUser(t *testing.T) {
	req, err := http.NewRequest("PUT", "/updateUser/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(usermngmnt.UpdateUser)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("UpdateUser handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}