/*
✅ Задача 10 — Редирект
Условие:
Создай сервер, который при обращении по адресу /redirect перенаправляет пользователя на сайт https://example.com.

🔧 Подсказка:

Используй http.Redirect(w, r, "https://example.com", http.StatusFound)

Код http.StatusFound (302) указывает на временный редирект.
*/

package gpt_tasks_http

import (
	"fmt"
	"net/http"
)

func redirHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://example.com", http.StatusFound)
}

func Task10() {
	http.HandleFunc("/redirect", redirHandler)

	fmt.Println("Server listening at :8080")
	http.ListenAndServe(":8080", nil)
}