/*
TODO:
4. Отправить GET-запрос из клиента
Напиши клиент, который делает GET-запрос на https://httpbin.org/get и печатает status code и body.

🔹 Подсказка: http.Get, io.ReadAll, defer resp.Body.Close()
*/

package gpt_tasks_http

import(
	"fmt"
	"io"
	"net/http"
)

func Task4(){
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Printf("Request error: %v", err)
		return 
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read request: %v", err)
		return
	}

	fmt.Println("Status: ", resp.Status)
	fmt.Println("Answer: ")
	fmt.Println(string(body))
}