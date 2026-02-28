package tests

import (
	"encoding/json"
	"net/http"
	"testing"
)

func BenchmarkGetTodos(b *testing.B) {
	getRouter()
	expected := []any{testTodo}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := benchReq(b, http.MethodGet, "/api/v1/todos", "")

		b.StopTimer()
		if w.Code != http.StatusOK {
			b.Fatalf("expected 200, got %d body=%s", w.Code, w.Body.String())
		}
		real := decodeJSON(b, w.Body.Bytes())
		assertJSONEqual(b, real, expected)
		b.StartTimer()
	}
}

func BenchmarkGetTodoByID(b *testing.B) {
	getRouter()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := benchReq(b, http.MethodGet, "/api/v1/todos/1", "")

		b.StopTimer()
		if w.Code != http.StatusOK {
			b.Fatalf("expected 200, got %d body=%s", w.Code, w.Body.String())
		}
		real := decodeJSON(b, w.Body.Bytes())
		assertJSONEqual(b, real, testTodo)
		b.StartTimer()
	}
}

func BenchmarkPostTodo(b *testing.B) {
	getRouter()
	payloadBytes, _ := json.Marshal(testTodo)
	payload := string(payloadBytes)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := benchReq(b, http.MethodPost, "/api/v1/todos", payload)

		b.StopTimer()
		if w.Code != http.StatusCreated {
			b.Fatalf("expected 201, got %d body=%s", w.Code, w.Body.String())
		}
		real := decodeJSON(b, w.Body.Bytes())
		assertJSONEqual(b, real, testTodo)
		b.StartTimer()
	}
}

func BenchmarkPutTodo(b *testing.B) {
	getRouter()
	payloadBytes, _ := json.Marshal(testTodo)
	payload := string(payloadBytes)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := benchReq(b, http.MethodPut, "/api/v1/todos/1", payload)

		b.StopTimer()
		if w.Code != http.StatusOK {
			b.Fatalf("expected 200, got %d body=%s", w.Code, w.Body.String())
		}
		real := decodeJSON(b, w.Body.Bytes())
		assertJSONEqual(b, real, testTodo)
		b.StartTimer()
	}
}

func BenchmarkDeleteTodo(b *testing.B) {
	getRouter()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := benchReq(b, http.MethodDelete, "/api/v1/todos/1", "")

		b.StopTimer()
		if w.Code != http.StatusOK {
			b.Fatalf("expected 200, got %d body=%s", w.Code, w.Body.String())
		}
		real := decodeJSON(b, w.Body.Bytes())
		assertJSONEqual(b, real, testTodo)
		b.StartTimer()
	}
}
