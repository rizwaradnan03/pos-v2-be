package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseType string

const (
	Success ResponseType = "success"
	Fail    ResponseType = "fail"
)

func BaseResponse(ctx *gin.Context, data interface{}, message string, err error) {
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  Fail,
			"message": err.Error(),
		})
		return
	}

	if message == "" {
		message = "Success"
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  Success,
		"data":    data,
		"message": message,
	})
}
