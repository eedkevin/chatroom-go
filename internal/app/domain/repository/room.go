package repository

import (
	"chatroom-demo/internal/app/domain"
)

type IRoomRepo interface {
	List() ([]domain.Room, error)
	Get(roomID string) (domain.Room, error)
	Save(domain.Room) error
	Delete(roomID string) error
}
