# Mock CBR Service

Mock-сервис для получения курсов валют (аналог API ЦБ РФ) с детерминированной генерацией данных.

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

Технологии:

- Go (net/http)
- XML
- Swagger
- Redis

---

## Запуск Redis:

```bash
docker run -d -p 6379:6379 redis
```
или

```bash
make redis-up
```