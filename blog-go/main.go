package main

import (
	"fmt"
	"log"

	"blog-go/config"
	"blog-go/database"
	"blog-go/router"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.InitConfig("config.yaml"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	if err := database.InitRedis(); err != nil {
		log.Printf("Warning: Failed to connect redis: %v", err)
	}

	gin.SetMode(config.AppConfig.Server.Mode)
	r := router.SetupRouter()

	addr := fmt.Sprintf(":%d", config.AppConfig.Server.Port)
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
