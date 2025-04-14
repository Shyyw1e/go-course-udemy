/*
✅ Задача 8 — Обработка маршрутов с параметрами
Условие:
Реализуй сервер, который отвечает на путь:

bash
Копировать
/hello/{name}
Когда пользователь заходит, например, на /hello/Alex, сервер должен ответить:

Копировать
Hello, Alex!
🔧 Подсказка:
В net/http нет встроенного роутера с параметрами, поэтому ты можешь:

либо распарсить r.URL.Path вручную, например через strings.Split(...)

либо использовать сторонние роутеры (например, chi, gorilla/mux) — но сейчас мы тренируемся на чистом net/http

💡 План:
Обработчик регистрируется как http.HandleFunc("/hello/", handler)

Из r.URL.Path вырезается имя пользователя (всё, что после /hello/)

Возвращается приветствие
*/

package gpt_tasks_http

import (
	"fmt"
	"net/http"
	"strings"
)

func nameHandler(w http.ResponseWriter, r *http.Request) {
	urla := strings.Split(r.URL.Path, "/")
	name := urla[len(urla) - 1]
	
	fmt.Fprintf(w, "Hello, %s!", name)
}

func Task8() {
	http.HandleFunc("/hello/", nameHandler)

	fmt.Println("Server listening at :8080")
	http.ListenAndServe(":8080", nil)
}