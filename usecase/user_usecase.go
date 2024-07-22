package usecase

import (
	"clean_sample/model"
	"clean_sample/repository"
)

type UserUsecase interface {
	CreateUser(user *model.User) error
	GetUserByID(id uint) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
	CreateUsersBulk(users []model.User) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) CreateUser(user *model.User) error {
	return u.userRepository.CreateUser(user)
}

func (u *userUsecase) GetUserByID(id uint) (*model.User, error) {
	return u.userRepository.GetUserByID(id)
}

func (u *userUsecase) UpdateUser(user *model.User) error {
	return u.userRepository.UpdateUser(user)
}

func (u *userUsecase) DeleteUser(id uint) error {
	return u.userRepository.DeleteUser(id)
}

func (u *userUsecase) CreateUsersBulk(users []model.User) error {
	return u.userRepository.CreateUsersBulk(users)
}
