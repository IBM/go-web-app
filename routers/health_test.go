package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	router := gin.Default()
	router.GET("/health", HealthGET)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("You received a %v error.", w.Code)
	}

	expected := "{\"status\":\"UP\"}"
	actual := w.Body.String()

	if actual != expected {
		t.Errorf("Response should be %v, was %v.", expected, actual)
	}
}
