package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"mortgage-calculator/pkg/handler"
)

type Server struct {
	router      *gin.Engine
	bankHandler handler.BankHandler
}

func NewServer(router *gin.Engine, bankHandler handler.BankHandler) *Server {
	return &Server{
		router:      router,
		bankHandler: bankHandler,
	}
}

func (s *Server) Run() error {
	r := s.Routes()
	err := r.Run()
	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}
	return nil
}
