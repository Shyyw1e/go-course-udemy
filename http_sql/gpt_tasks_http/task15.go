/*
Задача
Добавь маршрут /static/, который отдаёт файлы из папки ./assets/. Например, запрос к /static/logo.png должен вернуть файл assets/logo.png.

Ключевая подсказка
Подумай, как объединить http.StripPrefix и http.FileServer так, чтобы URL-путь /static/… «отрезался» и дальше шёл прямой маппинг на файловую систему.
*/

package gpt_tasks_http

import (

	"net/http"
)

func Task15() {
	fs := http.FileServer(http.Dir("./assets"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}