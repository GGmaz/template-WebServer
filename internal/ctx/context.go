package ctx

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	Transaction string = "transaction"
)

type RequestContext struct {
	gin.Context
}

func (rc *RequestContext) GetTransaction() *gorm.DB {
	if transaction, ok := rc.Value(Transaction).(*gorm.DB); ok {
		return transaction
	}
	return nil
}
