package wire

import (
	"github.com/GGmaz/wallet/user-service/internal/db/model"
	"github.com/GGmaz/wallet/user-service/internal/repo"
	"github.com/GGmaz/wallet/user-service/internal/service"
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
