package repository

import (
	"fmt"
	"chatroom-demo/internal/app/domain"
	"chatroom-demo/internal/app/infrastructure"

	"github.com/pkg/errors"
)

type UserRepo struct {
	storage infrastructure.IStorage
}

func NewUserRepo(storage infrastructure.IStorage) *UserRepo {
	return &UserRepo{storage: storage}
}

func (r UserRepo) List() ([]domain.User, error) {
	data, err := r.storage.List()
	if err != nil {
		return []domain.User{}, errors.Wrap(err, "error on UserRepo.List")
	}

	users := make([]domain.User, 0)
	for _, d := range data {
		user, ok := d.(domain.User)
		if !ok {
			return []domain.User{}, fmt.Errorf("error on UserRepo.List, %v", d)
		}
		users = append(users, user)
	}
	return users, nil
}

func (r UserRepo) Get(userID string) (domain.User, error) {
	data, err := r.storage.Get(userID)
	if err != nil {
		return domain.User{}, errors.Wrap(err, "error on UserRepo.Get")
	}

	if data == nil {
		return domain.User{}, nil
	}

	user, ok := data.(domain.User)
	if !ok {
		return domain.User{}, fmt.Errorf("error on UserRepo.Get converting to user, %v", data)
	}
	return user, nil
}

func (r UserRepo) Save(user domain.User) error {
	err := r.storage.Save(user.ID, user)
	return errors.Wrap(err, fmt.Sprintf("error on UserRepo.Save, %v", user))
}

func (r UserRepo) Delete(userID string) error {
	err := r.storage.Delete(userID)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error on UserRepo.Delete, %v", userID))
	}
	return nil
}
