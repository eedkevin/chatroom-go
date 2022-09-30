package service

import "chatroom-demo/internal/app/domain"

type IUserService interface {
	Get(userID string) (*domain.User, error)
	Create(userName string) (*domain.User, error)
	Delete(userID string) error
}
