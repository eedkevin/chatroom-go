package service

import (
	"chatroom-demo/internal/app/domain"
	"chatroom-demo/internal/app/domain/repository"
	"chatroom-demo/internal/app/domain/vo"

	"github.com/pkg/errors"
)

type ChatRoomService struct {
	roomRepo repository.IRoomRepo
}

func NewChatRoomService(roomRepo repository.IRoomRepo) *ChatRoomService {
	return &ChatRoomService{roomRepo: roomRepo}
}

func (s ChatRoomService) Get(roomID string) (domain.Room, error) {
	room, err := s.roomRepo.Get(roomID)
	if err != nil {
		return domain.Room{}, errors.Wrap(err, "error on getting room from storage")
	}

	return room, nil
}

func (s ChatRoomService) Create(roomName, roomType string) (domain.Room, error) {
	room, err := domain.NewRoom(roomName, roomType)
	if err != nil {
		return *room, err
	}
	err = s.roomRepo.Save(*room)
	if err != nil {
		return domain.Room{}, err
	}

	return *room, nil
}

func (s ChatRoomService) List() ([]domain.Room, error) {
	return s.roomRepo.List()
}

func (s ChatRoomService) Destroy(roomID string) error {
	return s.roomRepo.Delete(roomID)
}

func (s ChatRoomService) Join() error {
	return nil
}

func (s ChatRoomService) Thumbnail() (vo.RoomThumbnail, error) {
	return vo.RoomThumbnail{}, nil
}

func (s ChatRoomService) Participants() ([]domain.User, error) {
	return make([]domain.User, 0), nil
}
