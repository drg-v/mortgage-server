package app

import "github.com/gin-gonic/gin"

func (s *Server) Routes() *gin.Engine {
	router := s.router
	banks := router.Group("/banks")
	{
		banks.GET("/:id", s.bankHandler.GetBank())
		banks.GET("", s.bankHandler.GetAllBanks())
		banks.POST("", s.bankHandler.CreateBank())
		banks.PUT("", s.bankHandler.UpdateBank())
		banks.DELETE("/:id", s.bankHandler.DeleteBank())
		banks.POST("/mortgage", s.bankHandler.CalculateMortgage())
	}
	return router
}
