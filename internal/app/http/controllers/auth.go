package controllers

import (
	"pos-v2-be/internal/app/dtos"
	serviceinterface "pos-v2-be/internal/app/interfaces/service_interface"
	"pos-v2-be/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service  serviceinterface.AuthServiceInterface
	response func(*gin.Context, interface{}, string, error)
}

func NewAuthController(service serviceinterface.AuthServiceInterface) *AuthController {
	return &AuthController{service: service, response: response.BaseResponse}
}

func (ctrl *AuthController) SignIn(c *gin.Context) {
	req := &dtos.AuthSignInDto{}
	c.ShouldBindJSON(req)

	process, err := ctrl.service.SignIn(*req)

	ctrl.response(c, process, "Berhasil Melakukan Login", err)
}

func (ctrl *AuthController) TestDto(c *gin.Context) {
	req := &dtos.AuthSignInDto{}
	c.ShouldBindJSON(req)

	process, err := ctrl.service.SignIn(*req)

	ctrl.response(c, process, "Berhasil Melakukan Login", err)
}

func (ctrl *AuthController) RefreshAccessToken(c *gin.Context) {
	req := &dtos.RefreshAccessToken{}
	c.ShouldBindJSON(req)

	process, err := ctrl.service.RefreshAccessToken(*req)

	ctrl.response(c, process, "Berhasil Melakukan Login", err)
}

func (ctrl *AuthController) SignUp(c *gin.Context) {
	req := &dtos.AuthSignInUpDto{}
	c.ShouldBindJSON(req)

	process, err := ctrl.service.SignUp(*req)

	ctrl.response(c, process, "Berhasil Melakukan Login", err)
}