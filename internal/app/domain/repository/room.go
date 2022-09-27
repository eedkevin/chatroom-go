package repository

import (
	"chatroom-demo/internal/app/domain"
	"chatroom-demo/internal/app/domain/vo"
)

type IRoomRepo interface {
	List() ([]domain.Room, error)
	Get(roomID string) (domain.Room, error)
	Save(domain.Room) error
	Delete(roomID string) error
	Update(room domain.Room) error
	SaveMessage(roomID string, message vo.Message) error
	ListMessages(roomID string) ([]vo.Message, error)
}
