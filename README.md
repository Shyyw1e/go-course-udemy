# Конспект по пакету net/http в Go

Пакет `net/http` — стандартный способ работы с HTTP в Go. Он позволяет как **создавать HTTP-клиенты**, так и **HTTP-серверы**.

---

## 1. HTTP-Клиент (отправка запросов)

### Быстрый GET-запрос
```go
resp, err := http.Get("https://example.com")
```

### POST-запрос
```go
body := strings.NewReader("name=John&age=30")
resp, err := http.Post("https://example.com/form", "application/x-www-form-urlencoded", body)
```

### Использование `http.NewRequest` (гибкий способ)
```go
req, err := http.NewRequest("GET", "https://example.com", nil)
req.Header.Set("Authorization", "Bearer token")

client := &http.Client{}
resp, err := client.Do(req)
```

### Чтение ответа
```go
defer resp.Body.Close()
bodyBytes, err := io.ReadAll(resp.Body)
fmt.Println(string(bodyBytes))
```

---

## 2. HTTP-Сервер (обработка входящих запросов)

### Простой сервер
```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    log.Println("Server started on :8080")
    http.ListenAndServe(":8080", nil)
}
```

### Разделение маршрутов (ручное)
```go
http.HandleFunc("/hello", helloHandler)
http.HandleFunc("/ping", pingHandler)
```

### Чтение query-параметров
```go
r.URL.Query().Get("name")
```

### Чтение тела POST-запроса
```go
defer r.Body.Close()
body, _ := io.ReadAll(r.Body)
```

---

## 3. Полезные типы и методы

- `http.Request` — входящий запрос (метод, URL, заголовки, тело)
- `http.ResponseWriter` — запись ответа клиенту
- `http.Client` — отправка исходящих запросов
- `http.NewRequest(method, url, body)` — создание запроса вручную
- `http.ListenAndServe(addr, handler)` — запуск сервера

---

## 4. Обработка JSON

### Клиент: чтение JSON-ответа
```go
var data map[string]interface{}
json.NewDecoder(resp.Body).Decode(&data)
```

### Сервер: приём JSON и ответ
```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    var u User
    json.NewDecoder(r.Body).Decode(&u)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(u)
}
```

---

## 5. Советы

- Используй `http.Client` с `Timeout`
- Не забывай закрывать `resp.Body`
- Добавляй заголовки явно (`req.Header.Set`)
- Используй логгирование (`log.Println`) для отладки
- Лучше создавать отдельный хендлер на каждый маршрут

---

Хочешь — добавлю раздел о middleware или построении REST API.
