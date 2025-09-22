package handlers

import (
    "fmt"
    "html/template"
    "net/http"
    "net/url"
    "time"
    "url-shortener/internal/models"
    "url-shortener/internal/storage"
)

// Handlers - структура для обработчиков
type Handlers struct {
    store storage.Storage
}

// NewHandlers - создает новые обработчики
func NewHandlers(store storage.Storage) *Handlers {
    return &Handlers{store: store}
}

// Home - главная страница
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    
    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    tmpl.Execute(w, nil)
}

// Shorten - обработчик для сокращения URL
func (h *Handlers) Shorten(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    longURL := r.FormValue("url")
    if longURL == "" {
        http.Error(w, "URL is required", http.StatusBadRequest)
        return
    }
    
    // Проверяем валидность URL
    if _, err := url.ParseRequestURI(longURL); err != nil {
        http.Error(w, "Invalid URL", http.StatusBadRequest)
        return
    }
    
    // Генерируем короткий код (пока просто пример)
    shortCode := generateShortCode()
    
    // Сохраняем в хранилище
    urlMapping := &models.URLMapping{
        ShortURL:   shortCode,
        LongURL:    longURL,
        CreatedAt:  time.Now(),
        ClickCount: 0,
    }
    
    h.store.Save(urlMapping)
    
    // Возвращаем результат
    shortURL := fmt.Sprintf("http://localhost:8080/r/%s", shortCode)
    fmt.Fprintf(w, "Short URL: <a href='%s'>%s</a>", shortURL, shortURL)
}

// Redirect - обработчик для редиректа
func (h *Handlers) Redirect(w http.ResponseWriter, r *http.Request) {
    shortCode := r.URL.Path[len("/r/"):]
    
    if shortCode == "" {
        http.Error(w, "Short code required", http.StatusBadRequest)
        return
    }
    
    urlMapping, err := h.store.Find(shortCode)
    if err != nil || urlMapping == nil {
        http.NotFound(w, r)
        return
    }
    
    // Редирект на оригинальный URL
    http.Redirect(w, r, urlMapping.LongURL, http.StatusFound)
}

// generateShortCode - генерирует короткий код (упрощенная версия)
func generateShortCode() string {
    // Временная реализация - позже улучшим
    return "abc123"
}