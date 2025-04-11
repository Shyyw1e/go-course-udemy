/*
5. Пример с POST-запросом
Клиент отправляет JSON на https://httpbin.org/post и выводит JSON-ответ сервера.

🔹 Подсказка: http.NewRequest("POST", ...) + req.Header.Set("Content-Type", "application/json")
*/

package gpt_tasks_http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func Task5() {
	jsondata := []byte(`{"name": "Alex", "age": 30}`)
	req, err := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewReader(jsondata))
	if err != nil {
		fmt.Printf("invalid request: %s", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("invalid client: %s", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read response: %s", err)
		return
	}

	fmt.Println("Response:")
	fmt.Println("Status: ", resp.Status)
	fmt.Println("Body: ")
	fmt.Println(string(body))


}