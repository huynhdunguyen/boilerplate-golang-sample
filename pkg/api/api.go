package api

import (
	v1 "github.com/huynhdunguyen/boilerplate-golang-sample/pkg/api/v1"
	"github.com/huynhdunguyen/boilerplate-golang-sample/pkg/utl/config"
	"github.com/huynhdunguyen/boilerplate-golang-sample/pkg/utl/logger"
	"github.com/huynhdunguyen/boilerplate-golang-sample/pkg/utl/server"
)

// Start is func start api
func Start(cfg *config.Config) error {
	// create instance
	s := server.New(cfg)

	// init all config
	l := logger.New(cfg.Log, s)
	g := s.Group("/api")
	g.GET("/healthcheck", v1.HealthCheck)
	server.Start(s, cfg, l)
	return nil
}
