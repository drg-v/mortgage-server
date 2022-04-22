package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mortgage-calculator/pkg/dto"
	"mortgage-calculator/pkg/service"
	"net/http"
	"strconv"
)

type BankHandler interface {
	CreateBank() gin.HandlerFunc
	GetBank() gin.HandlerFunc
	GetAllBanks() gin.HandlerFunc
	UpdateBank() gin.HandlerFunc
	DeleteBank() gin.HandlerFunc
	CalculateMortgage() gin.HandlerFunc
}

type bankHandler struct {
	bankService service.BankService
}

func NewBankHandler(bankService service.BankService) BankHandler {
	return &bankHandler{
		bankService: bankService,
	}
}

func (s *bankHandler) CreateBank() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var bank dto.BankDto
		err := c.ShouldBindJSON(&bank)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		err = s.bankService.Create(bank)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		response := map[string]string{
			"status": "success",
			"data":   "new bank created",
		}
		c.JSON(http.StatusOK, response)
	}
}

func (s *bankHandler) GetBank() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Printf("handler error(incorrect bank id): %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		bank, err := s.bankService.Get(id)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		c.JSON(http.StatusOK, bank)
	}
}

func (s *bankHandler) GetAllBanks() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		banks, err := s.bankService.GetAll()
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		c.JSON(http.StatusOK, banks)
	}
}

func (s *bankHandler) UpdateBank() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var bank dto.BankDto
		err := c.ShouldBindJSON(&bank)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		err = s.bankService.Update(bank)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		response := map[string]string{
			"status": "success",
			"data":   "bank updated",
		}
		c.JSON(http.StatusOK, response)
	}
}

func (s *bankHandler) DeleteBank() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Printf("handler error(incorrect bank id): %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		err = s.bankService.Delete(id)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		response := map[string]string{
			"status": "success",
			"data":   "bank deleted",
		}
		c.JSON(http.StatusOK, response)
	}
}

func (s *bankHandler) CalculateMortgage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var mortgage dto.MortgageDto
		err := c.ShouldBindJSON(&mortgage)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		monthlyPayment, err := s.bankService.CalculateMortgage(mortgage)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", err))
			return
		}
		response := map[string]string{
			"status":         "success",
			"monthlyPayment": fmt.Sprintf("%f", monthlyPayment),
		}
		c.JSON(http.StatusOK, response)
	}
}
