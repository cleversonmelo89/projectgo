package tests

import (
	. "../router"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRepoGitByUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/repo/git/user", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(GetReposByUser)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
