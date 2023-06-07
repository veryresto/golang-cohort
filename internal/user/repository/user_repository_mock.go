package user

import (
	entity "online-course/internal/user/entity"
	response "online-course/pkg/response"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

// CountTotalUser implements UserRepository
func (repository *UserRepositoryMock) TotalCountUser() int64 {
	panic("unimplemented")
}

// FindAll implements UserRepository
func (repository *UserRepositoryMock) FindAll(offset int, limit int) []entity.User {
	panic("unimplemented")
}

// Delete implements UserRepository
func (repository *UserRepositoryMock) Delete(entity entity.User) *response.Error {
	panic("unimplemented")
}

// FindOneById implements UserRepository
func (repository *UserRepositoryMock) FindOneById(id int) (*entity.User, *response.Error) {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil, nil
	}

	user := arguments.Get(0).(entity.User)

	return &user, nil
}

// Update implements UserRepository
func (repository *UserRepositoryMock) Update(entity entity.User) (*entity.User, *response.Error) {
	panic("unimplemented")
}

// FindOneByCodeVerified implements UserRepository
func (repository *UserRepositoryMock) FindOneByCodeVerified(codeVerified string) (*entity.User, *response.Error) {
	arguments := repository.Mock.Called(codeVerified)

	if arguments.Get(0) == nil {
		return nil, nil
	}

	user := arguments.Get(0).(entity.User)

	return &user, nil
}

// Create implements UserRepository
func (repository *UserRepositoryMock) Create(entity entity.User) (*entity.User, *response.Error) {
	panic("unimplemented")
}

// FindByEmail implements UserRepository
func (repository *UserRepositoryMock) FindByEmail(email string) (*entity.User, *response.Error) {
	panic("unimplemented")
}
