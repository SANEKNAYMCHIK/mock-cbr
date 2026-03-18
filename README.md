# Mock CBR Service

Mock-сервис для имитации API ЦБ РФ с детерминированной генерацией данных.

---

## Технологии:

- Go (net/http)
- XML
- Swagger
- Redis
- Docker-compose
- Makefile

---

## Запуск:
```bash
make run
```
или
```go
go run cmd/mock-cbr/main.go
```
Сервис доступен по адресу: http://localhost:8080

---

## API

`GET /rate`

**Query параметры:**

*test_id (string, обязательный)* — влияет на генерацию данных \
*status (int, необязательный)* — HTTP статус ответа

**Пример:**

`GET /rate?test_id=123&status=200`

---

Пример ответа:
```xml
<ValCurs Date="18.03.2026" name="Foreign Currency Market">
    <Valute ID="R01235">
        <NumCode>840</NumCode>
        <CharCode>USD</CharCode>
        <Nominal>1</Nominal>
        <Name>Доллар США</Name>
        <Value>78,1234</Value>
        <VunitRate>78,1234</VunitRate>
    </Valute>
</ValCurs>
```
---

## Swagger:

Запустить генерацию документации для API проекта
```bash
make swagger
```

Swagger находится по адресу: http://localhost:8080/swagger/index.html

---

**Особенности:**

- Детерминированная генерация через test_id
- Реалистичные значения (±10%)
- Повторяемый набор валют
- Поддержка кастомного HTTP статуса

---

## Redis:

Сервис использует Redis как опциональный слой кеширования.

Если **Redis недоступен:**
- сервис продолжает работать
- данные генерируются без кеша

Если **Redis доступен:**
- результаты кешируются по ключу test_id
- TTL кеша: 5 минут

### Запуск Redis:

```bash
docker run -d -p 6379:6379 redis
```
или

```bash
make redis-up
```

---

## Логирование:
Для удобства просмотра, был добавлен middleware для логирования запросов с выводом информации в удобном формате

Сервис логирует:
- входящие HTTP запросы
- метод, путь, IP клиента
- время выполнения запроса

Пример:
```bash
Started GET /rate from 127.0.0.1:54321
Completed GET /rate in 1.2ms
```

---

## Дополнительно:

- Детерминированная генерация через test_id (один test_id → одинаковый результат)
- Случайный, но воспроизводимый набор валют
- Реалистичные значения (±10% от базового курса)
- Поддержка различных HTTP статусов (200 / 500)
