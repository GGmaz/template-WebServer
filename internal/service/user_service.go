package service

import (
	"github.com/GGmaz/wallet/user-service/internal/db/model"
	"github.com/GGmaz/wallet/user-service/internal/repo"
	"github.com/gin-gonic/gin"
)

type UserServiceImpl struct {
	UserRepo repo.Repo[model.User]
}

func (s *UserServiceImpl) CreateUser(c *gin.Context, email string) (*model.User, error) {
	user := &model.User{
		Email: email,
	}
	dbRes := s.UserRepo.Create(c, user)
	if dbRes.Error != nil {
		return nil, dbRes.Error
	}

	return user, nil
}
