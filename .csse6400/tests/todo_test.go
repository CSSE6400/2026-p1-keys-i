package tests

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

var testTodo = map[string]any{
	"id":          float64(1), // json numbers decode as float64 into map[string]any
	"title":       "Watch CSSE6400 Lecture",
	"description": "Watch the CSSE6400 lecture on ECHO360 for week 1",
	"completed":   true,
	"deadline_at": "2026-02-27T18:00:00",
	"created_at":  "2026-02-20T14:00:00",
	"updated_at":  "2026-02-20T14:00:00",
}

type TB interface {
	Helper()
	Fatalf(format string, args ...any)
}

func decodeJSON(t TB, body []byte) any {
	t.Helper()
	var v any
	if err := json.Unmarshal(body, &v); err != nil {
		t.Fatalf("invalid json: %v body=%s", err, string(body))
	}
	return v
}

func assertJSONEqual(t TB, real any, expec any) {
	t.Helper()
	if !reflect.DeepEqual(real, expec) {
		rbyte, _ := json.Marshal(real)
		ebyte, _ := json.Marshal(expec)
		t.Fatalf("json mismatch\nreal: %s\nexpected: %s", rbyte, ebyte)
	}
}

func TestGetTodos(t *testing.T) {
	req := setup(t)

	w := req(http.MethodGet, "/api/v1/todos", "")
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d body=%s", w.Code, w.Body.String())
	}

	real := decodeJSON(t, w.Body.Bytes())
	expec := []any{testTodo}
	assertJSONEqual(t, real, expec)
}

func TestGetTodoByID(t *testing.T) {
	req := setup(t)

	w := req(http.MethodGet, "/api/v1/todos/1", "")
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d body=%s", w.Code, w.Body.String())
	}

	real := decodeJSON(t, w.Body.Bytes())
	assertJSONEqual(t, real, testTodo)
}

func TestPostTodo(t *testing.T) {
	req := setup(t)

	payloadBytes, _ := json.Marshal(testTodo)

	w := req(http.MethodPost, "/api/v1/todos", string(payloadBytes))
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d body=%s", w.Code, w.Body.String())
	}

	real := decodeJSON(t, w.Body.Bytes())
	assertJSONEqual(t, real, testTodo)
}

func TestPutTodo(t *testing.T) {
	req := setup(t)

	payloadBytes, _ := json.Marshal(testTodo)

	w := req(http.MethodPut, "/api/v1/todos/1", string(payloadBytes))
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d body=%s", w.Code, w.Body.String())
	}

	real := decodeJSON(t, w.Body.Bytes())
	assertJSONEqual(t, real, testTodo)
}

func TestDeleteTodo(t *testing.T) {
	req := setup(t)

	w := req(http.MethodDelete, "/api/v1/todos/1", "")
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d body=%s", w.Code, w.Body.String())
	}

	real := decodeJSON(t, w.Body.Bytes())
	assertJSONEqual(t, real, testTodo)
}
