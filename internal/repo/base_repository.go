package repo

import (
	"errors"
	"github.com/GGmaz/wallet/user-service/internal/ctx"
	"github.com/GGmaz/wallet/user-service/internal/db/model"
	"github.com/GGmaz/wallet/user-service/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

// TODO: clean up
type Repo[T model.BaseInterface] struct {
}

func (r Repo[T]) Create(g *gin.Context, t *T) *gorm.DB {
	database, ok := g.Get(ctx.Transaction)
	if !ok {
		utils.Handle(errors.New("transaction not initalized"))
	}
	db := database.(*gorm.DB)
	log.Printf("%v", t)
	return db.Create(&t)
}
