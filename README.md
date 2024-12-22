# Калькулятор

## Описание

Этот проект реализует веб-сервис, которые вычисляет значение математического выражения, переданного в HTTP-запросе.

## Установка и запуск

Для запуска проекта выполните следующие шаги:

1. Установите go
2. Склонируйте репозиторий:

```bash
git clone "https://github.com/arbuzick57/calc_go.git"
```

3. Перейдите в папку с проектом

```bash
cd calc_go
```

4. Запустите сервер

```bash
go run cmd/main.go
```

По умолчанию сервер запустится на порту `8080`. Если необходимо изменить порт, установите переменную окружения `PORT` перед запуском.

## Примеры запросов

1. **Успешный запрос**

### Тело запроса

```json
{
  "expression": "5+7"
}
```

### Ответ

**Статус-код: `200 OK`**

```json
{
  "result": "12"
}
```

2. **Некорректное выражение**

### Тело запроса

```json
{
  "expression": "52+a"
}
```

### Ответ

**Статус-код: `422 Unprocessable Entity`**

```json
{
  "error": "not numbers"
}
```

3. **Ошибка при вычислении**

### Тело запроса

```json
{
  "expression": "5/0"
}
```

### Ответ

**Статус-код: `500 Internal Server Error`**

```json
{
  "error": "division by zero is forbidden"
}
```
