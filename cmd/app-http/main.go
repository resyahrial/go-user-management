package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/resyahrial/go-template/config"
	route "github.com/resyahrial/go-template/internal/api/routes"
	"github.com/resyahrial/go-template/internal/api/server"
	"github.com/resyahrial/go-template/internal/repositories/pg"
	"github.com/resyahrial/go-template/pkg/graceful"
)

type (
	Flag struct {
		Environment string
	}
)

var (
	appFlag Flag
)

func init() {
	flag.StringVar(
		&appFlag.Environment,
		"env",
		"dev",
		"env of deployment, will load the respective yml conf file.",
	)
}

func main() {
	flag.Parse()
	config.InitConfig(appFlag.Environment)

	_ = pg.InitDatabase(config.GlobalConfig)

	serverEngine := server.InitGinEngine(config.GlobalConfig)
	if serverEngine == nil {
		log.Fatal("server failed to initialized")
	}

	log.Printf("Running http server on port : %v", config.GlobalConfig.App.ServerAppPort)
	graceful.RunHttpServer(context.Background(), &http.Server{
		Addr:    fmt.Sprintf(":%v", config.GlobalConfig.App.ServerAppPort),
		Handler: route.InitRoutes(serverEngine, pg.DbInstance),
	}, 10*time.Second)
}
