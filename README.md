# Система транзакций

При запросе на сервер создается транзакция для определенного клиента, по каждому клиенту выстраивается очередь. При обработке транзакций проверяется баланс клиента и изменяется денежная сумма.

# Содержание

1. [Запуск](#Запуск)
2. [API](#API)
3. [Реализация](#Реализация)

# Запуск

```
make run
```

or

```
docker-compose up
```

# GET /api/clients/
Получение списка клиентов в системе

Запрос:

```
curl --location --request GET 'http://localhost:8080/api/clients/'
```

Ответ:

```
{
    "data": [
        {
            "id": 1,
            "name": "pavel",
            "balance": 1000
        },
        {
            "id": 2,
            "name": "ivan",
            "balance": 2000
        }
    ]
}
```

## GET /api/clients/:id
Получение одного клиента по id

Запрос:

```
curl --location --request GET 'http://localhost:8080/api/clients/1'
```

Ответ:

```
{
    "data": {
        "id": 1,
        "name": "pavel",
        "balance": 1000
    }
}
```

## POST /api/clients/
Создание нового клиента

Запрос:

```
curl --location --request POST 'http://localhost:8080/api/clients/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "ivan",
    "balance": 2000
}'
```

Ответ:

```
{
    "data": {
        "id": 2,
        "name": "ivan",
        "balance": 2000
    }
}
```

## DELETE /api/clients/:id
Удаление клиента по id

Запрос:

```
curl --location --request DELETE 'http://localhost:8080/api/clients/1'
```

Ответ:

```
Код ответа 200
```

## GET /api/transactions
Получение списка всех транзакций в системе

Запрос:

```
curl --location --request GET 'http://localhost:8080/api/transactions'
```

Ответ:

```
{
    "data": [
        {
            "id": 1,
            "senderId": 1,
            "receiverId": 2,
            "amount": 500,
            "status": "done",
            "updatedAt": "2022-11-11T20:19:24.560107Z"
        },
        {
            "id": 2,
            "senderId": 2,
            "receiverId": 1,
            "amount": 200,
            "status": "created",
            "updatedAt": "2022-11-11T20:19:23.709949Z"
        }
    ]
}
```

## GET /api/transactions/:id
Получение транзакции по id

Запрос:

```
curl --location --request GET 'http://localhost:8080/api/transactions/1'
```

Ответ:

```
{
    "data": {
        "id": 1,
        "senderId": 1,
        "receiverId": 2,
        "amount": 500,
        "status": "done",
        "updatedAt": "2022-11-11T20:19:24.560107Z"
    }
}
```

## POST /api/transactions/
Создание новой транзакции
- senderId - id отправителя
- receiverId - id получателя
- amount - сумма перевода

Запрос:

```
curl --location --request POST 'http://localhost:8080/api/transactions/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "senderId": 2,
    "receiverId": 1,
    "amount": 200
}'
```

Ответ:

```
{
    "data": {
        "id": 2,
        "senderId": 2,
        "receiverId": 1,
        "amount": 200,
        "status": "created",
        "updatedAt": "2022-11-11T20:19:23.709949592Z"
    }
}
```

## POST /api/clients/:id/transactions
Получение списка транзакций конкретного клиента с id

Запрос:

```
curl --location --request GET 'http://localhost:8080/api/clients/1/transactions'
```

Ответ:

```
{
    "data": [
        {
            "id": 1,
            "senderId": 1,
            "receiverId": 2,
            "amount": 500,
            "status": "done",
            "updatedAt": "2022-11-11T20:19:24.560107Z"
        }
    ]
}
```

# Реализация

- В качестве базы данных используется PostgreSQL
- Приложение и БД запускаются в докере с переменными окружения, указанными в docker-compose.yml
- При падении сервера информация о транзакциях сохраняется
- Для демонстрации накапливания транзакций в очереди установлена задержка в 10 секунд при обработке.
- Информация об успешности провередения транзакции логируется в терминал
- У каждого клиента своя собственная очередь, которая обрабатывается параллельно в горутинах
