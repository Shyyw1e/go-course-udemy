/*
✅ Задача 12 — Cookie и сессия
Условие:

Создай сервер, который:

При GET-запросе по пути /login устанавливает cookie с именем session_id и значением abc123.

При GET-запросе по пути /dashboard проверяет наличие cookie session_id.
Если cookie нет — возвращает 401 Unauthorized с сообщением "Please login first".
Если cookie есть и корректна — возвращает 200 OK с сообщением "Welcome to the dashboard".

📌 Подсказки:

Установка cookie:

go
Копировать
http.SetCookie(w, &http.Cookie{
    Name:  "session_id",
    Value: "abc123",
    Path:  "/",
})
Получение cookie:

go
Копировать
cookie, err := r.Cookie("session_id")
В случае ошибки можно проверить через errors.Is(err, http.ErrNoCookie).


*/

package gpt_tasks_http

import (
	"errors"
	"fmt"
	"net/http"
)

func cookieHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		http.SetCookie(w, &http.Cookie{
			Name: "session_id",
			Value: "abc123",
			Path: "/",
		})
		fmt.Fprintln(w, "Logged in")
	case "/dashboard":
		cookie, err := r.Cookie("session_id")
		if err != nil || errors.Is(err, http.ErrNoCookie) {
			http.Error(w, "Please login first", http.StatusUnauthorized)
			return
		}
		if cookie.Value != "abc123" {
			http.Error(w, "Invalid Session", http.StatusUnauthorized)
			return
		}
		fmt.Fprintln(w, "Welcome to dashboard")
	default:
		http.NotFound(w, r)
	}
}


func Task12() {
	http.HandleFunc("/", cookieHandler)

	http.ListenAndServe(":8080", nil)
}