package websockets

import (
	"pos-v2-be/internal/initial/intf"

	"github.com/gin-gonic/gin"
)

func Run(r *gin.Engine, service *intf.Services) {
	Ws(r, service)
}
