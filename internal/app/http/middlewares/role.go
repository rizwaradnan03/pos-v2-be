package middlewares

import (
	"pos-v2-be/internal/enums"
	"pos-v2-be/internal/pkg/auth"
	"slices"

	"github.com/gin-gonic/gin"
)

// func RoleMiddleware(allowedRoute []enums.RoleType) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		decodedToken, err := pkg.ExtractJwtAndDecode(c)
// 		if err != nil {
// 			response.BaseResponse(c, "", "Tidak Mendapatkan Akses!", validate.NewError("Gagal Verifikasi Token!"))

// 			c.Abort()
// 			return
// 		}

// 		isValid := slices.Contains((allowedRoute), decodedToken.Role)
// 		if !isValid {
// 			response.BaseResponse(c, "", "Tidak Mendapatkan Akses!", validate.NewError("Gagal Verifikasi Token!"))

// 			c.Abort()
// 			return
// 		}

// 		c.Next()
// 	}
// }

func RoleMiddleware(allowedRoute []enums.RoleType) gin.HandlerFunc {
	return func(c *gin.Context) {
		decodedToken, err := auth.ExtractJwtAndDecode(c)
		if err != nil {
			c.JSON(401, gin.H{
				"message": "Tidak Mendapatkan Akses!",
				"error":   "Gagal Verifikasi Token!",
			})
			c.Abort()
			return
		}

		isValid := slices.Contains(allowedRoute, decodedToken.Role)
		if !isValid {
			c.JSON(401, gin.H{
				"message": "Tidak Mendapatkan Akses!",
				"error":   "Gagal Verifikasi Token!",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
