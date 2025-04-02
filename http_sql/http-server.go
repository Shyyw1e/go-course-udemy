package httpsql

import (
	"encoding/json"
	"log"
	"net/http"
)

var (
	users = []User{{1, "Vasya"}, {2, "Tanya"}}
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resp)

}

func Serv() {
	http.HandleFunc("/users", handleUsers)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	
}