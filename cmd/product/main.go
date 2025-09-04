package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/jbakhtin/marketplace-cart/pkg/closer"
	"github.com/jbakhtin/marketplace-cart/pkg/starter"
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/config"
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/logger/zap"
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/server/rest"
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/storage/postgres"
	"github.com/jbakhtin/marketplace-product/internal/infrastructure/storage/postgres/repositories"
	"github.com/jbakhtin/marketplace-product/internal/modules/product"
)

var err error
var logger zap.Logger
var str starter.Starter
var clr closer.Closer
var cfg config.Config
var db *sql.DB
var restServer rest.Server

func init() {
	cfg, err = config.NewConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	logger, err = zap.NewLogger(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Инициализация клиента БД
	db, err = postgres.NewSQLClient(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	starterBuilder := starter.New()
	closerBuilder := closer.New()

	// Добавляем закрытие БД в closer
	closerBuilder.Add(func(ctx context.Context) error {
		return db.Close()
	})

	// Создаем репозиторий с подключением к БД
	productStorage := repositories.NewProductStorage(db)

	productModule, err := product.InitModule(logger, productStorage)
	if err != nil {
		log.Fatal(err)
	}

	restServer, err = rest.NewWebServer(&cfg, logger, productModule)
	if err != nil {
		log.Fatal(err)
	}
	starterBuilder.Add(restServer.Start)
	closerBuilder.Add(restServer.Shutdown)

	str = starterBuilder.Build()
	clr = closerBuilder.Build()
}

func main() {
	osCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cancel()

	if err = str.Start(osCtx); err != nil {
		logger.Error(err.Error())
		return
	}

	<-osCtx.Done()
	reason := "shutdown"
	if e := osCtx.Err(); e != nil {
		reason = e.Error()
	}
	logger.Info("received stop signal", reason)

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()
	if err = clr.Close(shutdownCtx); err != nil {
		logger.Error(err.Error())
	}
}
