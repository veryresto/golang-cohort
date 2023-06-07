package user

import (
	"testing"

	entity "online-course/internal/user/entity"
	repository "online-course/internal/user/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &repository.UserRepositoryMock{Mock: mock.Mock{}}
var userTestUsecase = NewUserUseCase(userRepository)

func TestUserUsecase_FindByIdSuccess(t *testing.T) {
	userData := entity.User{
		ID:    1,
		Name:  "faerulsalamun",
		Email: "faerulsalamun@gmail.com",
	}

	userRepository.Mock.On("FindOneById", 1).Return(userData)

	user, err := userTestUsecase.FindOneById(1)

	assert.NotNil(t, user)
	assert.Nil(t, err)
}

func TestUserUsecase_FindOneByCodeVerifiedSuccess(t *testing.T) {
	userData := entity.User{
		ID:    1,
		Name:  "faerulsalamun",
		Email: "faerulsalamun@gmail.com",
	}

	userRepository.Mock.On("FindOneByCodeVerified", "12345").Return(userData)

	user, err := userTestUsecase.FindOneByCodeVerified("12345")

	assert.NotNil(t, user)
	assert.Nil(t, err)
}

func TestUserUsecase_FindOneByCodeVerifiedNotFound(t *testing.T) {
	userRepository.Mock.On("FindOneByCodeVerified", "1234").Return(nil)

	user, err := userTestUsecase.FindOneByCodeVerified("1234")

	assert.Nil(t, user)
	assert.Nil(t, err)
}

func TestUserUsecase_FindByIdNotFound(t *testing.T) {
	userRepository.Mock.On("FindOneById", 2).Return(nil)

	user, err := userTestUsecase.FindOneById(2)

	assert.Nil(t, user)
	assert.Nil(t, err)
}
