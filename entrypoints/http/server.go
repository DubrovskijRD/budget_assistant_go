package http

import (
	"fmt"
	"log/slog"

	"github.com/DubrovskijRD/budget_assistant_go/application"
	"github.com/DubrovskijRD/budget_assistant_go/application/services"
	"github.com/DubrovskijRD/budget_assistant_go/entrypoints/http/controllers"
	"github.com/DubrovskijRD/budget_assistant_go/infrastructure"
	"github.com/DubrovskijRD/budget_assistant_go/infrastructure/models"
	"github.com/DubrovskijRD/budget_assistant_go/infrastructure/repositories"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Start() error
	ShutDown()
}

type ServerImpl struct {
	db     infrastructure.Database
	log    *slog.Logger
	router *gin.Engine
	cfg    application.Config
}

func NewServer(g *gin.Engine, db infrastructure.Database, log *slog.Logger, cfg application.Config) Server {
	return &ServerImpl{
		db:     db,
		router: g,
		log:    log,
		cfg:    cfg,
	}
}

func (s *ServerImpl) SetupRoutes() {

	repo := repositories.NewReceiptRepo(*s.db.GetDb())
	rs := services.NewReceiptService(repo)
	rc := controllers.NewReceiptController(rs, s.log)
	s.router.GET("/b/:id/labels", rc.AddReceipt)
	s.router.GET("/b/:id/receipts", rc.GetReceipts)
	s.router.POST("/b/:id/receipts", rc.AddReceipt)
}

func (s *ServerImpl) Start() error {
	s.InitDb()
	s.SetupRoutes()
	return s.router.Run(fmt.Sprintf(":%s", s.cfg.ServerPort))
}

func (s *ServerImpl) ShutDown() {
	s.db.Close()
}

func (s *ServerImpl) InitDb() {
	s.db.GetDb().AutoMigrate(&models.ReceiptItem{})
	s.db.GetDb().AutoMigrate(&models.ReceiptLabel{})
	s.db.GetDb().AutoMigrate(&models.Receipt{})
}
