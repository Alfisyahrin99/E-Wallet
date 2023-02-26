package usecase

import (
	"do-app/model"
	"do-app/repository"
	"do-app/utils"
)

type UserUseCase interface {
	ViewAllUser(page int, totalRows int) ([]model.User, error)
	CreateNewUser(newUser *model.User)error
	UpdateUser(user model.User) error
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func (u *userUseCase) ViewAllUser(page int, totalRows int) ([]model.User, error) {
	return u.userRepo.ViewAll(page, totalRows)
}

func (u *userUseCase) CreateNewUser(newUser *model.User)error {
	newUser.Id = utils.GenerateId()
	return u.userRepo.CreateNew(newUser)
}

func (u *userUseCase) UpdateUser(user model.User) error {
	if len(user.Name) < 3 || len(user.Name) > 20 {
		return error.New("Nama Minimal 3 Sampai 20 karakter")
	} else if (user.PhoneNumber)
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepository,
	}
}