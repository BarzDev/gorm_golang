package delivery

import (
	"log"

	"github.com/gin-gonic/gin"
)


type Server struct {
	engine *gin.Engine
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

	// ROUTE
	engine := gin.Default()


	return &Server{engine: engine}
}