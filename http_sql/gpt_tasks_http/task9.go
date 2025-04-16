/*
9. HTTP-сервер с маршрутом и JSON-ответом
Условие:
Создай сервер, который на GET-запрос /status возвращает JSON с текущим временем в поле "time" и сообщением "ok" в поле "status".

📌 Подсказка:

Используй json.NewEncoder(w).Encode(...)

Не забудь задать заголовок Content-Type: application/json
*/

package gpt_tasks_http

import (
	"encoding/json"
	"net/http"
	"time"
)

type jsonOk struct {
	Time time.Time `json:"time"`
	Status string `json:"status"`
}

func jsonOkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed: ", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	v := jsonOk{
	Status : "ok",
	Time : time.Now(),
	}
	
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		http.Error(w, "Failed to encode json: ", http.StatusBadRequest)
	}
	
}


func Task9() {
	http.HandleFunc("/status", jsonOkHandler)

	http.ListenAndServe(":8080", nil)
}