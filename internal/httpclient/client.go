package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second, // Устанавливаем таймаут
		},
	}
}

// Метод для выполнения GET-запроса
func (c *Client) Get(url string) (*http.Response, error) {
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении GET-запроса: %w", err)
	}
	return resp, nil
}

// Метод для выполнения POST-запроса с JSON
func (c *Client) Post(url string, data interface{}) (*http.Response, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации данных: %w", err)
	}

	resp, err := c.httpClient.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении POST-запроса: %w", err)
	}
	return resp, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}
