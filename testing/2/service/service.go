package service

import (
	"github.com/fn-code/Go-Example/testing/2/storage"
)

type Service interface {
	RegisterUser(cmd *RegisterUser) (bool, error)
}

type ServiceImpl struct {
	userRepo storage.UserRepo
}

func NewServiceImpl(userRepo storage.UserRepo) *ServiceImpl {
	return &ServiceImpl{userRepo: userRepo}
}

func (s ServiceImpl) RegisterUser(cmd *RegisterUser) (bool, error) {

	userCheck, err := s.userRepo.FindByEmail(cmd.Email)
	if err != nil {
		return false, err
	}

	if userCheck == nil {
		return false, nil
	}

	userInfo := storage.UserInfo{
		ID:     "",
		Name:   cmd.Name,
		Email:  cmd.Email,
		Gender: cmd.Gender,
	}
	err = s.userRepo.Save(userInfo)
	if err != nil {
		return false, err
	}
	return true, nil
}
