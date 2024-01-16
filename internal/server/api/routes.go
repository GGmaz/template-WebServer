package v1

import (
	v1 "github.com/GGmaz/wallet/user-service/internal/server/api/v1"
	"github.com/gin-gonic/gin"
)

func RegisterVersion(router *gin.Engine) {
	v1.RegisterUser(router)
}
