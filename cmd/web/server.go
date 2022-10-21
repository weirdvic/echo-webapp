package main

import (
	"context"
	"flag"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/weirdvic/echo-tutorial/pkg/models/postgres"
	"go.uber.org/zap"
)

type application struct {
	db_dsn        *string
	echo          *echo.Echo
	http_endpoint *string
	logger        *zap.SugaredLogger
	snippets      *postgres.SnippetModel
}

func main() {
	// Create middleware logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	db_dsn := flag.String("dsn", "postgres://snippetbox:snippetbox@localhost:5432/snippetbox?sslmode=disable", "Database URL")
	http_endpoint := flag.String("addr", ":1323", "HTTP server address")
	flag.Parse()

	db, err := pgxpool.New(context.Background(), *db_dsn)
	if err != nil {
		logger.Sugar().Fatal("Unable to create connection pool: %v", err)
	}
	defer db.Close()

	// Create application instance
	app := &application{
		db_dsn:        db_dsn,
		echo:          echo.New(),
		http_endpoint: http_endpoint,
		logger:        logger.Sugar(),
		snippets: &postgres.SnippetModel{
			DB: db,
		},
	}

	// Attach logger for requests
	app.echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	}))

	// Init server instance
	app.init()

	// Parse command line flags and start the server

	app.echo.Logger.Fatal(app.echo.Start(*app.http_endpoint))
}
