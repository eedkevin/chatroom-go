package repository

import (
	"chatroom-demo/internal/app/domain"
	"chatroom-demo/internal/app/domain/vo"
	"chatroom-demo/internal/app/infrastructure"
	"fmt"

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

func (r RoomRepo) Get(roomID string) (*domain.Room, error) {
	data, err := r.storage.Get(roomID)
	if data == nil && err == nil { // not found
		return nil, fmt.Errorf("NOT_FOUND")
	}

	if err != nil {
		return nil, errors.Wrap(err, "error on RoomRepo.Get")
	}

	room, ok := data.(domain.Room)
	if !ok {
		return nil, fmt.Errorf("error on RoomRepo.Get converting to room, %v", data)
	}
	return &room, nil
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

func (r RoomRepo) Update(room domain.Room) error {
	data, err := r.storage.Update(room.ID, room)
	if data == nil && err == nil { // not found
		return fmt.Errorf("NOT_FOUND")
	}

	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error on RoomRepo.Update, %v", room))
	}
	return nil
}

func (r RoomRepo) SaveMessage(roomID string, message vo.Message) error {
	data, err := r.storage.Get(roomID)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error on RoomRepo.SaveMessage, %v, %v", roomID, message))
	}
	room, ok := data.(domain.Room)
	if !ok {
		return fmt.Errorf("error on RoomRepo.SaveMessage, %v", data)
	}
	room.Messages = append(room.Messages, message)
	_, err = r.storage.Update(room.ID, room)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error on RoomRepo.SaveMessage, %v, %v", roomID, message))
	}
	return nil
}

func (r RoomRepo) ListMessages(roomID string) ([]vo.Message, error) {
	data, err := r.storage.Get(roomID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error on RoomRepo.ListMessages, %v, %v", roomID, data))
	}
	room, ok := data.(domain.Room)
	if !ok {
		return nil, fmt.Errorf("error on RoomRepo.ListMessages, %v", data)
	}
	return room.Messages, nil
}
