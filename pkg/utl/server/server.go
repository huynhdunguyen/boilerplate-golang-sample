package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"github.com/huynhdunguyen/boilerplate-golang-sample/pkg/utl/config"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// New is func create new instance
func New(cfg *config.Config) *gin.Engine {
	gin.SetMode(cfg.Server.GinMode)

	e := gin.New()
	e.Static("/openapi", "./openapi")
	e.Use(gin.Recovery())

	return e
}

// Start is func start instance
func Start(e *gin.Engine, cfg *config.Config, l zap.Logger) {
	h := fmt.Sprintf("0.0.0.0:%s", cfg.Server.Port)

	srv := &http.Server{
		Addr:         h,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      e,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	l.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		l.Info("Server Shutdown")
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		l.Info("timeout of 5 seconds.")
	}
	l.Info("Server exiting")
}
