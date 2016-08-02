package middleware

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestFormChecker(t *testing.T) {
	form := url.Values{}
	form.Add("asd", "azina")
	form.Add("dani", "azuna")

	req, err := http.NewRequest("POST", "/test1", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := gin.Default()
	router.POST("/test1", FormChecker("ali", "dani"))

	router.ServeHTTP(rr, req)

	expected := `{"error":"invalid jwt"}`
	resp := strings.TrimSpace(rr.Body.String())

	if resp != expected {
		t.Error("Form Checker failed")
	}
}
