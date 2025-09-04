[![codecov](https://codecov.io/gh/jbakhtin/marketplace-product/graph/badge.svg?token=5REv1ymRVB)](https://codecov.io/gh/jbakhtin/marketplace-product)
## Marketplace Product Service

Коротко: сервис продуктов маркетплейса. Простые команды для запуска через Docker.

### 1) Переменные окружения

Скопируйте пример и при необходимости поменяйте значения:
```bash
cp env.example .env
```

Ключевые переменные:
- `WEBSERVER_RESTAPI_ADDRESS` — адрес HTTP-сервера (по умолчанию `:8080`)
- `DATABASE_DRIVER` — драйвер БД (по умолчанию `pgx`)
- `DATABASE_DSN` — строка подключения к Postgres

### 2) Команды Makefile

```bash
make build  # Сборка контейнеров
make start  # Запуск docker compose (в фоне)
make stop   # Остановка docker compose
```

После `make start` приложение поднимется, применит миграции и будет доступно на адресе из `WEBSERVER_RESTAPI_ADDRESS`.

### 3) Эндпоинты

- `GET  /healthz`  — проверка живости
- `GET  /readyz`   — готовность
- `GET  /product/get` — получить товар по SKU (см. хендлер)
- `POST /product/list` — получить список SKU (см. хендлер)

### Примечания

- Миграции применяются автоматически при старте (см. `internal/infrastructure/storage/postgres/storage.go`).
- Логи — через Zap.
- Роутер — Chi. Аутентификация применяется для группы `/product`.