package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	todo "github.com/CSSE6400/2026-p1-keys-i/internal/todo"
	"github.com/gin-gonic/gin"
)

var (
	once sync.Once
	rtr  *gin.Engine
)

func getRouter() *gin.Engine {
	once.Do(func() {
		gin.SetMode(gin.TestMode)
		rtr = todo.NewRouter()
	})
	return rtr
}

func setup(t *testing.T) func(method, path, body string) *httptest.ResponseRecorder {
	t.Helper()
	t.Parallel()

	return func(method, path, body string) *httptest.ResponseRecorder {
		var req *http.Request
		if body == "" {
			req = httptest.NewRequest(method, path, http.NoBody)
		} else {
			req = httptest.NewRequest(method, path, bytes.NewReader([]byte(body)))
			req.Header.Set("Content-Type", "application/json")
		}

		req.Header.Set("Accept", "application/json")

		w := httptest.NewRecorder()
		getRouter().ServeHTTP(w, req)
		return w
	}
}
