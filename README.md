# REST API Для менеджмента книг

## Используемые технологии:
- Go 1.18
- Postgres

## Дополнительно:
- Auth JWT. Middleware
- Clean architecture
- Gin Gonic
- Graceful Shutdown
- Swagger

### Для инициализации приложения:

```
make init
```

### Для запуска/остановки БД:

```
make db.start
```

```
make db.stop
```

### Для запуска приложения:

```
make run
```
### Для обновления Swagger документации:

```
make swag
```

#### После развертывания перейдите по адресу:
localhost:8000/swagger/index.html


Доступные эндпоинты:
- [POST] /auth/sign-up - Регистрация
- [POST] /auth/sign-in - Вход в систему
- [POST] /api/books - добавить книгу
- [GET] /api/books - получить список книг
- [GET] /api/books/user - получить книги пользователя (берется из JWT)
- [GET] /api/books/:id - получить книгу по ID
- [PUT] /api/books/:id - обновить книгу
- [DELETE] /api/books/:id - удалить книгу