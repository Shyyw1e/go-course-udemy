/*
Задача 7 — Обработка POST-формы
Условие:
Создай сервер с маршрутом /form, который обрабатывает POST-запрос с обычной HTML-формой.

Форма отправляет два поля:

text
Копировать
name=Alex
age=30
Сервер должен распарсить форму и вывести:

sql
Копировать
Hello, Alex! You are 30 years old.
🔧 Подсказка:
Используй r.ParseForm()

Доступ к полям: r.FormValue("name"), r.FormValue("age")
💡 План:
Настроить обработчик /form

Обрабатывать только POST-запросы

Распарсить форму

Вывести результат
*/

package gpt_tasks_http

import (
	"fmt"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed: ", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse: ", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "html form name: %s \n", r.FormValue("name"))
	fmt.Fprintf(w, "html form name: %s \n", r.FormValue("age"))
}
func Task7() {
	http.HandleFunc("/form", formHandler)
	http.ListenAndServe(":8080", nil)
	
}