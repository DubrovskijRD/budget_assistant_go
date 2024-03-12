package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/DubrovskijRD/budget_assistant_go/application"
	"github.com/DubrovskijRD/budget_assistant_go/entrypoints/http"
	"github.com/DubrovskijRD/budget_assistant_go/infrastructure"
	"github.com/gin-gonic/gin"
)

func run() error {
	cfg := application.NewConfig()
	g := gin.Default()
	log := slog.Default()
	db := infrastructure.NewPostgresDatabase(
		&cfg,
	)
	defer db.Close()
	serv := http.NewServer(g, db, log, cfg)
	defer serv.ShutDown()
	return serv.Start()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}
