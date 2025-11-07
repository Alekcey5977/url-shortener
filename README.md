# URL Shortener

Простой сервис сокращения ссылок на Go с использованием памяти для хранения URL.  
Пользователь вводит длинную ссылку, сервис возвращает короткую ссылку, которая редиректит на оригинальный URL.

## Стек технологий

- Go
- HTML/CSS для фронтенда
- Память (MemoryStorage) для хранения URL
- Конкурентность: `sync.RWMutex`
- HTTP/REST API

## Функционал

- Сокращение длинной ссылки через `/shorten` (POST)
- Редирект с короткой ссылки через `/r/{shortCode}`
- Главная страница с формой для ввода URL

## Структура проекта

url-shortener/\
├── cmd/\
│ ├── server/\
│ │ ├── main.go/ # точка входа\
├── internal/\
│ ├── handlers/ # HTTP обработчики\
│ ├── models/ # структура URLMapping\
│ └── storage/ # интерфейс и реализация MemoryStorage\
├── templates/\
│ └── index.html # фронтенд форма

## Установка и запуск

1. Клонируйте репозиторий:
```bash
git clone https://github.com/Alekcey5977/url-shortener
cd url-shortener
```

2. Запустите приложение:
```go
cd cmd/server
go run main.go
```

3. Откройте браузер и перейдите по адресу:
```
http://localhost:8080/
```

4. Введите длинную ссылку и нажмите Shorten — получите короткую ссылку.

## Примеры использования:

Длинная ссылка: `https://example.com/very/long/url`

Сокращённая ссылка: `http://localhost:8080/r/abc123` (редиректит на оригинал)
