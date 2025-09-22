package models

import "time"

// URLMapping - структура для хранения сопоставления URL
type URLMapping struct {
	ShortURL  string	`json:"short_url"`
	LongURL   string	`json:"long_url"`
	CreatedAt time.Time	`json:"created_at"`
	ClickCount int	`json:"click_count"`
}
