/*
Условие:

Сделай GET-запрос на:

arduino
Копировать
https://httpbin.org/json
Распарси JSON-ответ от сервера и выведи title из поля slideshow.

🔎 Пример ответа от сервера:
json
Копировать

	{
	  "slideshow": {
	    "author": "Yours Truly",
	    "date": "date of publication",
	    "title": "Sample Slide Show",
	    "slides": [...]
	  }
	}

🔧 План действий:
Сделать GET-запрос

Прочитать тело ответа (io.ReadAll)

# Объявить структуру со вложенным slideshow.title

# Распарсить с помощью json.Unmarshal

Вывести slideshow.title
*/
package gpt_tasks_http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Slideshow struct {
		Author string `json:"author"`
		Date string `json:"date"`
		Title string `json:"title"`
	}`json:"slideshow"`
}

func Task6() {
	resp, err := http.Get("https://httpbin.org/json")
	if err != nil {
		fmt.Printf("invalid request: %s", err)
		return
	} 
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read response: %s", err)
	}
	var respjson Response
	json.Unmarshal(body, &respjson)
	fmt.Printf("Slideshow title: %s", respjson.Slideshow.Title)
}

