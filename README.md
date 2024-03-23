# REST API MOVIES

Этот проект представляет собой REST API для управления информацией об актерах в киноиндустрии.

## Запуск приложения

### Системные требования

- Docker
- Docker Compose

### Установка и запуск

1. Склонируйте репозиторий:

```bash
git clone https://github.com/ваш-пользователь/rest-api-go.git
```

2. Перейдите в директорию проекта:
```bash
cd rest-api-movies
```

3. Запустите приложение с помощью Docker Compose:
```bash
docker-compose up --build
```

Приложение будет запущено на порту 8080.

## Использование
- Создание нового актера:
```bash
curl -H "username: myuser" -H "password: mypassword" -X POST -d '{"name": "Brad Pitt", "gender": "Male", "birth_date": "1963-12-18"}' -H "Content-Type: application/json" http://localhost:8080/actors
```

- Получение списка всех актеров:
```bash
curl -H "username: readonly_user" -H "password: password" http://localhost:8080/actors-list
```

## Структура проекта
- `cmd/app/main.go`: Основной файл приложения, инициализирующий HTTP-сервер и обработчики маршрутов.

- `internal/app`: Директория с основной логикой приложения.

- `internal/database`: Пакет для работы с базой данных PostgreSQL.

- `internal/models`: Модели данных приложения.

- `config/config.go`: Файл с конфигурацией приложения.