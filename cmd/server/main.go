package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"mortgage-calculator/pkg/app"
	"mortgage-calculator/pkg/entity"
	"mortgage-calculator/pkg/handler"
	"mortgage-calculator/pkg/repository"
	"mortgage-calculator/pkg/service"
	"mortgage-calculator/util"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\\n", err)
		os.Exit(1)
	}
}

func run() error {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	connection := config.DBUsername + ":" + config.DBPassword + "@tcp(" + config.ServerAddress
	connection += ")/" + config.DBName + "?charset=utf8"
	db, err := setupDatabase(connection)
	if err != nil {
		return err
	}
	bankRepository := repository.NewBankRepository(db)
	router := gin.Default()
	router.Use(cors.Default())
	bankService := service.NewBankService(bankRepository)
	bankHandler := handler.NewBankHandler(bankService)
	server := app.NewServer(router, bankHandler)
	err = server.Run()
	if err != nil {
		return err
	}
	return nil
}

func setupDatabase(connection string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&entity.Bank{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
