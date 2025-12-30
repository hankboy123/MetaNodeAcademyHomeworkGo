package main

import (
	"log"
	"sh-manage/config"
	"sh-manage/handlers"
	"sh-manage/middleware"
	"sh-manage/models"
	"sh-manage/services"
	"sh-manage/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 加载配置
	cfg := config.Load("")

	db, err := gorm.Open(mysql.Open(config.GetMySQLDSN(cfg)), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// 自动迁移
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService, []byte(cfg.JWT.Secret))

	r := gin.Default()

	r.Use(middleware.Logger())
	r.Use(middleware.CORS())

	r.GET("health", func(c *gin.Context) {
		utils.Success(c, gin.H{
			"status": "ok",
		})
	})

	public := r.Group("/api/v1")
	{
		public.POST("/users/register", userHandler.Register)
		public.POST("/users/login", userHandler.Login)
	}

	// 需要认证的路由
	protected := r.Group("/api/v1")
	protected.Use(middleware.Auth([]byte(cfg.JWT.Secret)))
	{
		protected.GET("/users/me", userHandler.GetProfile)
		protected.PUT("/users/me", userHandler.UpdateProfile)
	}

	addr := cfg.Server.Host + ":" + cfg.Server.Port
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
