package service

import (
	"chatroom-demo/internal/app/application"
	"chatroom-demo/internal/app/domain"
	"chatroom-demo/internal/app/domain/repository"

	"github.com/pkg/errors"
)

type UserService struct {
	userRepo repository.IUserRepo
}

func NewUserService(userRepo repository.IUserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s UserService) Get(userID string) (*domain.User, error) {
	user, err := s.userRepo.Get(userID)

	if err != nil {
		if err.Error() == application.NotFoundErr {
			return nil, nil
		}
		return nil, errors.Wrap(err, "error on getting user from storage")
	}

	return user, nil
}

func (s UserService) Create(userName string) (*domain.User, error) {
	user, err := domain.NewUser(userName)
	if err != nil {
		return nil, err
	}
	err = s.userRepo.Save(*user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s UserService) Delete(userID string) error {
	return s.userRepo.Delete(userID)
}
