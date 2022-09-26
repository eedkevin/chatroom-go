package repository

import (
	"fmt"
	"chatroom-demo/internal/app/domain"
	"chatroom-demo/internal/app/infrastructure"

	"github.com/pkg/errors"
)

type RoomRepo struct {
	storage infrastructure.IStorage
}

func NewRoomRepo(storage infrastructure.IStorage) *RoomRepo {
	return &RoomRepo{storage: storage}
}

func (r RoomRepo) List() ([]domain.Room, error) {
	data, err := r.storage.List()
	if err != nil {
		return []domain.Room{}, errors.Wrap(err, "error on RoomRepo.List")
	}

	rooms := make([]domain.Room, 0)
	for _, d := range data {
		room, ok := d.(domain.Room)
		if !ok {
			return []domain.Room{}, fmt.Errorf("error on RoomRepo.List, %v", d)
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}

func (r RoomRepo) Get(roomID string) (domain.Room, error) {
	data, err := r.storage.Get(roomID)
	if err != nil {
		return domain.Room{}, errors.Wrap(err, "error on RoomRepo.Get")
	}

	if data == nil {
		return domain.Room{}, nil
	}

	room, ok := data.(domain.Room)
	if !ok {
		return domain.Room{}, fmt.Errorf("error on RoomRepo.Get converting to room, %v", data)
	}
	return room, nil
}

func (r RoomRepo) Save(room domain.Room) error {
	err := r.storage.Save(room.ID, room)
	return errors.Wrap(err, fmt.Sprintf("error on RoomRepo.Save, %v", room))
}

func (r RoomRepo) Delete(roomID string) error {
	err := r.storage.Delete(roomID)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error on RoomRepo.Delete, %v", roomID))
	}
	return nil
}
