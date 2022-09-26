package domain

import (
	"fmt"
	"chatroom-demo/internal/app/domain/vo"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/teris-io/shortid"
)

var ROOM_TYPE_PUBLIC = "public"
var ROOM_TYPE_PRIVATE = "private"

type Room struct {
	ID        string
	Code      string
	Name      string
	Type      string
	Thumbnail vo.RoomThumbnail
}

func NewRoom(roomName, roomType string) (*Room, error) {
	code, err := shortid.Generate()
	if err != nil {
		return &Room{}, errors.Wrap(err, fmt.Sprintf("error on generating shortid for room[%s,%s]", roomName, roomType))
	}

	id := uuid.Must(uuid.NewV4(), nil).String()

	room := &Room{
		ID:   id,
		Name: roomName,
		Code: code,
		Thumbnail: vo.RoomThumbnail{
			RoomID:         id,
			RoomName:       roomName,
			RoomType:       roomType,
			RecentMessages: make([]vo.Message, 0),
		},
		Type: roomType,
	}

	return room, nil
}
