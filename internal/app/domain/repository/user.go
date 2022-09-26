package repository

import "chatroom-demo/internal/app/domain"

type IUserRepo interface {
	List() ([]domain.User, error)
	Get(userID string) (domain.User, error)
	Save(domain.User) error
	Delete(userID string) error
}
