package delivery

import (
	"fmt"
	"library-api/config"
	"log"

	"github.com/gin-gonic/gin"
)


type Server struct {
	engine *gin.Engine
	host string
}

func (s *Server) InitRoute() {
	rg := s.engine.Group("/")

	rg.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

}

func (s *Server) Run() {
	s.InitRoute()
	if err := s.engine.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}


func NewServer() *Server{
	cfg, err := config.NewConfig()
	if err !=nil{
		log.Fatalf("config error : %v", err)
	}
	db := config.ConnectDB()
	_ = db

	// ROUTE
	engine := gin.Default()
	host := fmt.Sprintf(":%s",cfg.ApiPort )


	return &Server{engine: engine,host: host}
}