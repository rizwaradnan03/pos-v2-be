package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "pos-v2-be/cmd/docs"
	"pos-v2-be/internal/app/routes"
	"pos-v2-be/internal/app/scheduler"
	"pos-v2-be/internal/db"
	"pos-v2-be/internal/initial"
	"pos-v2-be/internal/pkg/validate"
)

func main() {
	validate.InitValidator()
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// Global CORS with explicit methods and preflight support
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: false,
		MaxAge:           86400,
	}))

	uploads := r.Group("/uploads")
	{
		uploads.Static("/", "./uploads")
		uploads.OPTIONS("/*filepath", func(c *gin.Context) {
			c.Status(http.StatusNoContent)
		})
	}

	// DB
	db.ConnectDatabase()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Initial
	init := initial.NewInit(db.DB)

	// Routes
	routes.WebSocketRoute(r, init.Infra.Service)
	routes.RegisterRoute(r, init)

	// Scheduler
	scheduler.RunSchedules(db.DB)

	// Serve
	port := os.Getenv("APP_PORT")
	r.Run("0.0.0.0:" + port)
}
