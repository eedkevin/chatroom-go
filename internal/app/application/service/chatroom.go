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
	Join() error
	Thumbnail() (vo.RoomThumbnail, error)
	Participants() ([]domain.User, error)
}
