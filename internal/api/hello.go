package api

import (
	"encoding/json"
	"net/http"
)

// Пример структуры ответа
type Response struct {
	Message string `json:"message"`
}

// Пример обработчика для API эндпоинта
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
