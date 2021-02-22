package handler

import (
	"bytes"
	"encoding/json"
	"github.com/fn-code/Go-Example/testing/2/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type serviceMock struct {
	mock.Mock
}

func (s *serviceMock) RegisterUser(cmd *service.RegisterUser) (bool, error) {
	args := s.Called(cmd)
	return args.Bool(0), args.Error(1)
}

func serviceImple(t *testing.T) *serviceMock {
	svc := &serviceMock{}
	svc.On("RegisterUser", mock.MatchedBy(func(cmd *service.RegisterUser) bool {
		assert.NotEqual(t, nil, cmd)
		return true
	})).Return(true, nil)
	return svc
}

func TestHandler_RegisteUserPage(t *testing.T) {
	svc := serviceImple(t)
	handler := &Handler{svc: svc}

	params := &service.RegisterUser{
		Name:   "ludin Nento",
		Email:  "email@gmail.com",
		Gender: "male",
	}

	buf, _ := json.Marshal(params)
	bodys := bytes.NewBuffer(buf)

	r, _ := http.NewRequest("POST", "localhost:8080/user/register", bodys)
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.RegisteUserPage(w, r)
	resp := w.Result()
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

}
