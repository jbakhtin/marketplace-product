# Marketplace Product Service

Микросервис для управления товарами в маркетплейсе.

## Настройка проекта

### Переменные окружения

Скопируйте файл `env.example` в `.env` и настройте переменные:

```bash
cp env.example .env
```

### Запуск базы данных

```bash
docker-compose up -d
```

### Запуск приложения

#### С помощью Makefile (рекомендуется):
```bash
# Полная настройка проекта
make dev-setup

# Запуск в режиме разработки
make dev

# Или сборка и запуск
make run
```

#### С помощью Go напрямую:
```bash
go run cmd/product/main.go
```

### Доступные команды Makefile

```bash
make help          # Показать все доступные команды
make setup         # Настройка проекта (deps + docker + .env)
make dev-setup     # Полная настройка для разработки
make build         # Сборка приложения
make run           # Сборка и запуск
make dev           # Запуск в режиме разработки
make test          # Запуск тестов
make clean         # Очистка артефактов сборки
make docker-up     # Запуск Docker Compose
make docker-down   # Остановка Docker Compose
```

## Структура базы данных

Приложение автоматически выполнит миграции при запуске. Таблица `products` будет создана с полями:
- `id` - первичный ключ
- `sku` - артикул товара
- `name` - название товара  
- `price` - цена товара

## Архитектура

Проект использует Clean Architecture с разделением на слои:

- **Domain** - бизнес-логика и сущности
- **Use Case** - сценарии использования
- **Infrastructure** - внешние зависимости (БД, HTTP, логирование)
- **Ports** - интерфейсы для инверсии зависимостей

## API Endpoints

- `GET /product/{sku}` - получение товара по SKU
- `GET /product/list` - получение списка SKU товаров