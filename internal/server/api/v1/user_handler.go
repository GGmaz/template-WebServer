package v1

import (
	"github.com/GGmaz/template-service/internal/db/model"
	"github.com/GGmaz/template-service/pkg/utils"
	"github.com/GGmaz/template-service/pkg/wire"
	"github.com/gin-gonic/gin"
)

func RegisterUser(router *gin.Engine) {
	v1 := router.Group("/api/v1/user")
	{
		v1.POST("/", create)
		v1.GET("/:email/balance", getBalance)
	}
}

func create(c *gin.Context) {
	println("create..")
	var createuser model.User
	err := c.BindJSON(&createuser)
	utils.Handle(err)
	if createuser.Email == "" {
		c.PureJSON(400, gin.H{"error": "email is required"})
		return
	}

	user, err := wire.Svc.UserService.CreateUser(c, createuser.Email)
	if err != nil {
		c.PureJSON(500, gin.H{"error": err.Error()})
		return
	}
	c.PureJSON(200, user)

}

func getBalance(c *gin.Context) {
	println("balance")
}
