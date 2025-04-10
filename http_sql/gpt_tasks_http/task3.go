/*	
3. Принять JSON в теле POST-запроса
Сервер принимает JSON по пути /user:

json
Копировать
{ "name": "Alex", "age": 30 }
И возвращает тот же объект в ответе.

🔹 Подсказка: json.NewDecoder(r.Body).Decode(&struct{})
*/

package gpt_tasks_http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task3User struct {
	Name string `json:"name"`
	Age int `json:"age,omitempty"`

}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var U Task3User
	err := json.NewDecoder(r.Body).Decode(&U)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(U)	

}

func Task3() {
	http.HandleFunc("/user", jsonHandler)
	fmt.Println("Listening server at :8080")
	http.ListenAndServe(":8080", nil)
}