package gpt_tasks_http

import (
	"fmt"
	"net/http"
)

/*
	1. Простой HTTP-сервер
Создай HTTP-сервер на порту :8080, который на запрос по пути /ping отвечает pong.

🔹 Подсказка: используй http.HandleFunc и http.ListenAndServe

*/
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}
func Task1() {
	http.HandleFunc("/ping", handler)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}