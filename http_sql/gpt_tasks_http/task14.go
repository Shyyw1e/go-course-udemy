/*
✅ Задача 14 — Обработка query‑параметров
Условие:
Добавь в сервер новый маршрут /greet, который принимает GET‑запрос с параметром name, например:

GET /greet?name=Alice
и возвращает текстовый ответ:

Hello, Alice!
Если параметр name отсутствует или пуст, то возвращает:

Hello, stranger!

📌 Подсказка:

Проверь метод через r.Method == http.MethodGet.

Достань параметр через r.URL.Query().Get("name").

Если строка пуста — подставь "stranger".

Установи заголовок Content-Type: text/plain.

Отправь ответ через fmt.Fprintln(w, ...).
*/
package gpt_tasks_http

import (
	"fmt"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "stranger"
	}

	fmt.Fprintf(w, "Hello, %s\n", name)
}

func Task14() {
	http.HandleFunc("/greet", greetHandler)

	http.ListenAndServe(":8080", nil)
}
