package app

import (
	"item-detail-api/config"
	"item-detail-api/packages/dbconn"
	"item-detail-api/src/app/web"
	"item-detail-api/src/core/seed"
	"item-detail-api/src/infrastructure/dependencies"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Start() {
	conf, err := config.Get()
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	logger, err := initLogger(conf)
	if err != nil {
		log.Println("failed to init logger:", err)
		os.Exit(1)
	}
	defer func() {
		_ = logger.Sync()
	}()

	dbConnector := dbconn.NewSQLiteConnector(config.SQLiteConfig(conf))
	db, err := dbConnector.Connect()
	if err != nil {
		logger.Fatal("failed to connect to SQLite DB", zap.Error(err))
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	seed.Seed(db, logger)

	webHandlers := dependencies.Build(conf, logger, db)
	web.StartWebServer(conf, webHandlers, logger, db)
}

func initLogger(conf *config.Config) (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return cfg.Build()
}
