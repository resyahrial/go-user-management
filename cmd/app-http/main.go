package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/resyahrial/go-user-management/config"
	"github.com/resyahrial/go-user-management/internal/api/server"
	"github.com/resyahrial/go-user-management/internal/repositories/pg"
	"github.com/resyahrial/go-user-management/pkg/graceful"
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
		"example",
		"env of deployment, will load the respective yml conf file.",
	)
}

func main() {
	flag.Parse()
	config.InitConfig(appFlag.Environment)

	_ = pg.InitDatabase(config.GlobalConfig)

	serverEngine := server.InitGinEngine(config.GlobalConfig, pg.DbInstance)
	if serverEngine == nil {
		log.Fatal("server failed to initialized")
	}

	log.Printf("Running http server on port : %v", config.GlobalConfig.App.ServerAppPort)
	graceful.RunHttpServer(context.Background(), &http.Server{
		Addr:    fmt.Sprintf(":%v", config.GlobalConfig.App.ServerAppPort),
		Handler: serverEngine,
	}, 10*time.Second)
}
