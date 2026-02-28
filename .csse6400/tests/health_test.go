package tests

import (
	"encoding/json"
	"net/http"
	"testing"
)


func TestHealth(t *testing.T) {
	req := setup(t)

	writer := req(http.MethodGet, "/api/v1/health", "")
	if writer.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d body=%s", writer.Code, writer.Body.String())
	}

	var got map[string]any
	if err := json.Unmarshal(writer.Body.Bytes(), &got); err != nil {
		t.Fatalf("response is not valid json: %v body=%s", err, writer.Body.String())
	}

	if got["status"] != "ok" {
		t.Fatalf(`expected status "ok", got %v`, got["status"])
	}
}
