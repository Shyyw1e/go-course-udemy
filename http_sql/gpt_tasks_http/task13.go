/*
Задача 13 — Обработка PUT‑запроса
Условие:
Сделайте обработчик для пути /update, который принимает только PUT‑запросы. В теле запроса приходит JSON вида:

json
Копировать
Редактировать
{ "value": "new data" }
Необходимо распарсить поле value, сохранить его в переменную и вернуть клиенту текстовый ответ "Updated".

Подсказка:
Используйте:

проверку метода через r.Method == http.MethodPut,

defer r.Body.Close() после открытия тела,

json.NewDecoder(r.Body).Decode(&struct{Value string}{}) для разбора JSON,

установку статуса и заголовка ответа (w.WriteHeader, w.Header().Set("Content-Type", "text/plain")),

отправку ответа через fmt.Fprintln(w, "Updated").
*/

package gpt_tasks_http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type jsonParse struct {
	Value string `json:"value"`
	Date time.Time `json:"time,omitempty"`
}

var jsonSlice []jsonParse
func jsonPutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	v := jsonParse{}
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		http.Error(w, "Failed to decode json", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "text/plain")
	jsonSlice = append(jsonSlice, v)
	fmt.Fprintln(w, "Updated")
}

func Task13() {
	http.HandleFunc("/update", jsonPutHandler)
	fmt.Println("\n\nListening on :8080")
	http.ListenAndServe(":8080", nil)
}