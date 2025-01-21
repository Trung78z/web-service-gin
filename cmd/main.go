package main

import (
	"log"

	"github.com/Trung78z/web-service-gin/config"
	"github.com/Trung78z/web-service-gin/internal/repositories"
	"github.com/Trung78z/web-service-gin/internal/routes"
	"github.com/Trung78z/web-service-gin/pkg/cache"
	"github.com/Trung78z/web-service-gin/pkg/database"
)

func main() {
	cfg := config.Load()
	log.Println("App is running on port:", cfg.AppPort)
	log.Println("Connecting to database:", cfg.DBConn)
	log.Println("Development:", cfg.Development)
	database.InitDB(cfg.DBConn)
	repositories.InitRepository()
	cache.InitCache(cfg.RedisURL)
	r := routes.InitRoutes(cfg.Development)
	r.Run(":" + cfg.AppPort)
}
