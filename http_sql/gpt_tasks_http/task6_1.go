/*
6. Сервер с несколькими маршрутами
Сервер по /ping возвращает "pong", по /time — текущее время.
*/

package gpt_tasks_http

import (
	"fmt"
	"net/http"
	"time"
)

func pingTimeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/ping":
		fmt.Fprintln(w, "/pong")
	case "/time":
		now := time.Now().Format("h15:04:05")
		fmt.Fprintf(w, "Time: %s\n", string(now))
	default:
		fmt.Fprintln(w, "Not Found")
	}
}
func Task6_1() {
	http.HandleFunc("/", pingTimeHandler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)

}