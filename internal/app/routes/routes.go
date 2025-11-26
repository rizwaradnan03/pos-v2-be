package routes

import (
	"pos-v2-be/internal/app/http/middlewares"
	"pos-v2-be/internal/app/http/websockets"
	"pos-v2-be/internal/initial"
	"pos-v2-be/internal/initial/intf"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine, init *initial.Init) {
	r.Use(middlewares.CorsMiddleware())
	r.Use(middlewares.LogMiddleware())

	api := r.Group("/api")
	{
		AuthRoute(api, init)
	}

}

func WebSocketRoute(r *gin.Engine, service *intf.Services) {
	websockets.Run(r, service)
}
