package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"

	"testing"
)

var (
	benchReqMu    sync.Mutex
	benchReqCache = map[string]*http.Request{}
)

// benchReq runs a single request through the shared router.
func benchReq(b *testing.B, method, path, body string) *httptest.ResponseRecorder {
	b.Helper()

	var req *http.Request

	// Fast path: no body -> reuse request
	if body == "" {
		key := method + " " + path

		benchReqMu.Lock()
		req = benchReqCache[key]
		if req == nil {
			req = httptest.NewRequest(method, path, http.NoBody)
			req.Header.Set("Accept", "application/json")
			benchReqCache[key] = req
		}
		benchReqMu.Unlock()
	} else {
		// Body path: create a fresh request each time (body reader is stateful)
		req = httptest.NewRequest(method, path, strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
	}

	w := httptest.NewRecorder()
	getRouter().ServeHTTP(w, req)
	return w
}
