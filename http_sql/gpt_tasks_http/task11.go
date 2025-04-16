/*
✅ Задача 11 — Простая авторизация через заголовок
Условие:
Создай сервер с маршрутом /private, который обрабатывает только GET-запросы, и проверяет наличие заголовка:

makefile
Копировать
Authorization: secret-token
Если заголовок не передан или передан неверно, сервер должен вернуть 401 Unauthorized и сообщение "Unauthorized".

Если заголовок корректный — вернуть 200 OK и "Welcome to the private area".

🔧 Подсказка:
Заголовки читаются через r.Header.Get("Authorization").

Для ответа можно использовать http.Error или fmt.Fprintln(w, ...).

Установи статус кода вручную при необходимости (w.WriteHeader(...)).

🧠 План:
Напиши обработчик /private.

Внутри обработчика:

Проверь метод.

Проверь заголовок Authorization.

Верни соответствующий ответ.


*/

package gpt_tasks_http

import (
	"fmt"
	"net/http"
)

func privateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Authorization") != "secret-token" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Fprintln(w, "Welcome to the private area")
}

func Task11() {
	http.HandleFunc("/private", privateHandler)

	http.ListenAndServe(":8080", nil)
}