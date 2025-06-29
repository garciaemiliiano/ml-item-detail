package web

import (
	"io"
	"item-detail-api/config"
	"item-detail-api/src/app/router"
	"item-detail-api/src/infrastructure/dependencies"
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func StartWebServer(conf *config.Config, handlers *dependencies.WebHandlerContainer, logger *zap.Logger, db *gorm.DB) {
	newRouter := router.Create()

	configureMappings(newRouter, handlers, conf, logger, db)
	s := &http.Server{
		Addr:         ":" + conf.PORT,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	s.Handler = http.TimeoutHandler(newRouter, time.Second*120, `{"message":"timeout"}`)
	s.ErrorLog = log.New(newLoggerWriter(logger), "", 0)

	if err := newRouter.Run(":" + config.GetConfig().PORT); err != nil {
		panic(err)
	}
}

type zapWriter struct {
	logger *zap.Logger
}

func (fw *zapWriter) Write(p []byte) (n int, err error) {
	fw.logger.Error(string(p))
	return len(p), nil
}

func newLoggerWriter(logger *zap.Logger) io.Writer {
	return &zapWriter{logger: logger}
}
