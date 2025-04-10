/*
2. Обработка query-параметров
На GET-запрос по /hello?name=Иван верни Привет, Иван!.

🔹 Подсказка: r.URL.Query().Get("name")



*/

package gpt_tasks_http

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "guest"
	}
	fmt.Fprintf(w, "Привет, %s", name)
}

func Task2() {
 	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}