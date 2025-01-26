package main

import (
	"os"
	"os/signal"
	"syscall"

	golog "github.com/Vladroon22/GoLog"
	"github.com/Vladroon22/TestTask-Bank-Operation/internal/database"
	"github.com/Vladroon22/TestTask-Bank-Operation/internal/handlers"
	"github.com/Vladroon22/TestTask-Bank-Operation/internal/repository"
	"github.com/Vladroon22/TestTask-Bank-Operation/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger := golog.New()

	if err := godotenv.Load(); err != nil {
		logger.Fatalln("Error loading .env file: ", err)
	}

	conn, err := database.DBConn()
	if err != nil {
		logger.Fatalln(err)
	}
	logger.Infoln("Database's connection is valid")

	repo := repository.NewRepository(conn, logger)
	srv := service.NewService(repo)
	handler := handlers.NewHandler(srv, logger)

	r := gin.Default()
	r.Use(gin.Recovery(), gin.Logger())

	r.POST("/api/v1/up", handler.IncreaseUserBalance)
	r.POST("/api/v1/transfer", handler.TransferMoney)
	r.GET("/api/v1/tx/:userID", handler.GetLastTxs)

	go func() {
		if err := r.Run(":8080"); err != nil {
			logger.Fatalln("Error start server: ", err)
		}
	}()

	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, syscall.SIGTERM, syscall.SIGINT)
	<-exitCh

	go func() {
		conn.Close()
	}()

	logger.Infoln("Gracefull shutdown")
}
