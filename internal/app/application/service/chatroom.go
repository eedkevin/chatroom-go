package service

import (
	"chatroom-demo/internal/app/domain"
	"chatroom-demo/internal/app/domain/vo"
)

type IChatRoom interface {
	Get(roomID string) (domain.Room, error)
	Create(roomName string, roomType string) (domain.Room, error)
	List() ([]domain.Room, error)
	Destroy(roomID string) error
	SaveMessage(roomID string, message vo.Message) error
	ListMessages(roomID string) ([]vo.Message, error)
	Join(roomID string) error
	Thumbnail(roomID string) (vo.RoomThumbnail, error)
	Participants(roomID string) ([]domain.User, error)
}
