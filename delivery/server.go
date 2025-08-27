package delivery

import (
	"fmt"
	"log"

	"library-api/config"
	"library-api/delivery/controller"
	"library-api/repository"
	"library-api/usecase"

	"github.com/gin-gonic/gin"
)

type Server struct {
	authorUC usecase.AuthorUseCase
	engine   *gin.Engine
	host     string
}

func (s *Server) InitRoute() {
	rg := s.engine.Group("/")
	controller.NewAuthorController(s.authorUC, rg).Route()
}

func (s *Server) Run() {
	s.InitRoute()
	if err := s.engine.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error : %v", err)
	}
	db := config.ConnectDB()
	// Inject DB ke -> Repository
	authorRepository := repository.NewAuthorRepository(db)

	// Inject Repository ke -> Usecase
	authorUC := usecase.NewAuthorUsecase(authorRepository)

	// ROUTE
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		authorUC: authorUC,
		engine:   engine,
		host:     host,
	}
}
