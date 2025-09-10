[![codecov](https://codecov.io/gh/jbakhtin/marketplace-product/graph/badge.svg?token=5REv1ymRVB)](https://codecov.io/gh/jbakhtin/marketplace-product)
## Marketplace Product Service

Коротко: сервис продуктов маркетплейса. Простые команды для запуска через Docker.

### 1) Переменные окружения

Скопируйте пример и при необходимости поменяйте значения:
```bash
cp env.example .env
```

Ключевые переменные:
```
# Web Server
WEBSERVER_REST_ADDRESS= #адрес HTTP-сервера (по умолчанию `:8080`)

# Database
DB_DRIVER=              #драйвер БД (по умолчанию `pgx`)
DB_HOST=                #для локальной разработки устанавливаем имя контейнера БД (по умолчанию 
`marketplace_product_db`)
DB_USER=                #имя пользователя для подключения к БД
DB_PASSWORD=            #пароль пользователя для подключения к БД
DB_NAME=                #имя БД
```

### 2) Команды Makefile

```bash
make build  # Сборка контейнеров
make start  # Запуск docker compose (в фоне)
make stop   # Остановка docker compose
make test   # Запуск тестов
```

После `make start` приложение поднимется, применит миграции и будет доступно на адресе из 
`WEBSERVER_REST_ADDRESS`.

### 3) Эндпоинты

- `GET  /healthz`  — проверка живости
- `GET  /readyz`   — готовность
- `GET  /products/get` — получить товар по SKU
- `GET  /products/list` — получить список SKU

### Примечания
- Миграции применяются автоматически при старте;
- Логи — через Zap;
- Роутер — Chi;