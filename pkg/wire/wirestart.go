package wire

import (
	"github.com/GGmaz/template-service/internal/db/model"
	"github.com/GGmaz/template-service/internal/repo"
	"github.com/GGmaz/template-service/internal/service"
)

type Wires struct {
	UserService service.UserServiceImpl
}

func Init() *Wires {
	w := Wires{
		UserService: service.UserServiceImpl{
			UserRepo: repo.Repo[model.User]{},
			//UserCompanyRepo: repo.Repo[model.UserCompany]{},
		},
	}
	return &w
}

var Svc = Init()
