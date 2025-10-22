package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"testovoe/internal/config"
	"testovoe/internal/handler"
	"testovoe/internal/repository"
	"testovoe/internal/service"
)

// @title Test Task
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
func main() {
	config.LoadEnv()

	pool := repository.Connect()
	repo := repository.NewRepository(pool)
	serv := service.NewService(repo)
	h := handler.NewHandler(serv)

	r := gin.Default()
	r.Use(cors.Default())
	h.Router(r)

	fmt.Printf("Starting server on port: %s\n", config.Env.ServerPort)
	err := r.Run(":" + config.Env.ServerPort)
	if err != nil {
		log.Fatalf("server start failed: %v", err)
	}
}
