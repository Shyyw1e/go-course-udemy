/*
### Задача: Простой REST-эндпоинт для управления заметками

**Условие:**

Реализуй два маршрута для работы с «заметками» в памяти:

1. **GET /api/notes**
   ­ Возвращает список всех заметок в виде JSON-массива объектов

   ```json
   [
     { "id": 1, "text": "First note" },
     { "id": 2, "text": "Another note" }
   ]
   ```
2. **POST /api/notes**
   ­ Принимает в теле JSON с полем `text`, создаёт новую заметку с уникальным целым `id` и тем же полем `text`, добавляет её в память и возвращает HTTP 201 с полным объектом заметки в теле ответа.

**Требования:**

* Храни заметки в глобальном слайсе или мапе, защищённом `sync.Mutex`.
* При любом GET `/api/notes` возвращай актуальный срез.
* При POST:

  * Проверяй в `r.Method == http.MethodPost`, иначе — 405.
  * Парсь JSON через `json.NewDecoder(r.Body).Decode(&req)` и не забывай `defer r.Body.Close()`.
  * Если в запросе нет поля `text` или оно пустое — возвращай 400 Bad Request.
  * Генерируй новый `id` (можно просто увеличивать счётчик).
  * Устанавливай заголовки:

    ```go
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    ```
  * Отправляй ответ через `json.NewEncoder(w).Encode(newNote)`.

**Подсказки:**

* Используй `http.HandleFunc("/api/notes", notesHandler)`.
* Внутри `notesHandler` разветвляй логику по методу (`switch r.Method { case "GET": ... case "POST": ... }`).
* Для конкурентного доступа к слайсу/мапе заметок — `var mu sync.Mutex`.

---
*/

package gpt_tasks_http

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Note struct{
	ID int `json:"id"`
	Text string `jspn:"text,omitempty"`
}

var(
	notes []Note
	nextID = 1
	mu sync.Mutex
)

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")

		mu.Lock()
	
		data := make([]Note, len(notes))
		copy(data, notes)
		mu.Unlock()

		json.NewEncoder(w).Encode(data)
		
	case "POST":
		defer r.Body.Close()
	
		var req struct {
			Text string `json:"text"`
		}
	
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		if req.Text == "" {
			http.Error(w, "Empty JSON text", http.StatusBadRequest)
			return 
		}
		mu.Lock()
		note := Note{ID: nextID, Text: req.Text}
		nextID++
		notes = append(notes, note)
		mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(note)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Task16() {
	http.HandleFunc("/api/notes", NotesHandler)

	http.ListenAndServe(":8080", nil)
	
}