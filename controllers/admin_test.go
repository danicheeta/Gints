package controllers

/*import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var rr *httptest.ResponseRecorder
var ai *AdminImpl

func init() {
	router = gin.Default()
	rr = httptest.NewRecorder()
	ai = &AdminImpl{Router: router, DB: nil, Generator: nil, AuthWare: nil}
}

func TestSetGameConf(t *testing.T) {
	var jsonStr = []byte(`{"name":"ttl", "geners": ["gener"], "banner": "bnr"}`)

	req, err := http.NewRequest("PUT", "/SetGameConf", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	router.PUT("/SetGameConf", ai.SetGameConf)
	router.ServeHTTP(rr, req)

	//expected := GameConf{Name: "ttl", Geners: "geners", Banner: "bnr"}
}
*/
