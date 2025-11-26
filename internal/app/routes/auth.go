package routes

import (
	"pos-v2-be/internal/initial"

	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.RouterGroup, init *initial.Init) {

	rg := r.Group("/auth")
	{
		rg.POST("/sign-in", init.Infra.Controller.AuthController.SignIn)
		rg.POST("/refresh-access-token", init.Infra.Controller.AuthController.RefreshAccessToken)
		rg.POST("/sign-up", init.Infra.Controller.AuthController.SignUp)
	}
}
