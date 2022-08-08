# CURRENCY SERVICE

### Table "currency_pair"
| currency_from | currency_to | well |       updated_at       |
|---------------|:-----------:|:----:|:----------------------:|
| USD           |     RUB     |  75  | 2020-09-23 09:13:00+00 |


### API:
```
1. Создание записи

POST /api/currency
{
    "currencyFrom": "USD",
    "currencyTo": "RUB"
}

2. Перевод значения из одной валюты в другую

PUT /api/currency
{
    "currencyFrom": "USD",
    "currencyTo": "RUB"
    "value":  1
}

3. Агрегация добавленных валютных пар

GET  /api/currency
[
    {
        "currencyFrom": "USD",
        "currencyTo": "RUB"
        ….
    }
]