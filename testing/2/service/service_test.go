package service

import (
	"github.com/fn-code/Go-Example/testing/2/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type UserRepoMock struct {
	mock.Mock
}

func (u UserRepoMock) FindByEmail(email string) (*storage.UserInfo, error) {
	args := u.Called(email)
	return args.Get(0).(*storage.UserInfo), args.Error(1)
}

func (u UserRepoMock) Save(info storage.UserInfo) error {
	args := u.Called(info)
	return args.Error(0)
}

func TestServiceImpl_RegisterUser(t *testing.T) {
	userRepo := &UserRepoMock{}
	service := &ServiceImpl{userRepo: userRepo}

	cmd := &RegisterUser{
		Name:   "User 1",
		Gender: "Male",
		Email:  "male@gmail.com",
	}

	userRepo.On("Save", mock.MatchedBy(func(req storage.UserInfo) bool {
		return true
	})).Return(nil)

	userRepo.On("FindByEmail", mock.MatchedBy(func(email string) bool {
		test := "male1@gmail.com"
		assert.NotEqual(t, test, email)
		return true
	})).Return(&storage.UserInfo{}, nil)

	ok, err := service.RegisterUser(cmd)
	assert.NoError(t, err)
	assert.Equal(t, true, ok)
}
