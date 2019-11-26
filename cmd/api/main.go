package main

import (
	"flag"
	"github.com/huynhdunguyen/boilerplate-golang-sample/pkg/api"
	"github.com/huynhdunguyen/boilerplate-golang-sample/pkg/utl/config"
)

// @title Source Go API docs
// @version 1.0
// @description This is a sample server celler server.
// @host localhost:8080
// @BasePath /api

func main() {
	cfgPath := flag.String("p", "./cmd/api/config.local.yaml", "Path to config file")
	flag.Parse()
	cfg, err := config.Load(*cfgPath)
	checkErr(err)
	checkErr(api.Start(cfg))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
