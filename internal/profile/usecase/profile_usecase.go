package profile

import (
	oauthUsecase "online-course/internal/oauth/usecase"
	dto "online-course/internal/profile/dto"
	userDto "online-course/internal/user/dto"
	userEntity "online-course/internal/user/entity"
	userUsecase "online-course/internal/user/usecase"
	"online-course/pkg/response"
)

type ProfileUsecase interface {
	FindProfile(id int) (*dto.ProfileResponseBody, *response.Error)
	Update(id int, dto userDto.UserUpdateRequestBody) (*userEntity.User, *response.Error)
	Deactive(id int) *response.Error
	Logout(accessToken string) *response.Error
}

type profileUsecase struct {
	userUsecase  userUsecase.UserUsecase
	oauthUsecase oauthUsecase.OauthUsecase
}

// Deactive implements ProfileUsecase
func (usecase *profileUsecase) Deactive(id int) *response.Error {
	// Get profile
	user, err := usecase.userUsecase.FindOneById(id)

	if err != nil {
		return err
	}

	err = usecase.userUsecase.Delete(int(user.ID))

	if err != nil {
		return err
	}

	return nil
}

// Logout implements ProfileUsecase
func (usecase *profileUsecase) Logout(accessToken string) *response.Error {
	return usecase.oauthUsecase.Logout(accessToken)
}

// Update implements ProfileUsecase
func (usecase *profileUsecase) Update(id int, dto userDto.UserUpdateRequestBody) (*userEntity.User, *response.Error) {
	// Get profile
	user, err := usecase.userUsecase.FindOneById(id)

	if err != nil {
		return nil, err
	}

	updateUser, err := usecase.userUsecase.Update(int(user.ID), dto)

	if err != nil {
		return nil, err
	}

	return updateUser, nil
}

// FindProfile implements ProfileUsecase
func (usecase *profileUsecase) FindProfile(id int) (*dto.ProfileResponseBody, *response.Error) {
	// Get profile
	user, err := usecase.userUsecase.FindOneById(id)

	if err != nil {
		return nil, err
	}

	userResponse := dto.CreateProfileResponse(*user)

	return &userResponse, nil
}

func NewProfileUsecase(
	userUsecase userUsecase.UserUsecase,
	oauthUsecase oauthUsecase.OauthUsecase,
) ProfileUsecase {
	return &profileUsecase{
		userUsecase,
		oauthUsecase,
	}
}
