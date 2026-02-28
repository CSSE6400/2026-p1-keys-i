package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func BenchmarkHealth(b *testing.B) {
	router := getRouter()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/health", http.NoBody)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			b.Fatalf("expected 200, got %d body=%s", w.Code, w.Body.String())
		}

		// cheap correctness check (no json unmarshal)
		if strings.TrimSpace(w.Body.String()) != `{"status":"ok"}` {
			b.Fatalf("unexpected body=%s", w.Body.String())
		}
	}
}
