package domain

import (
	"chatroom-demo/internal/app/domain/vo"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/teris-io/shortid"
)

var ROOM_TYPE_PUBLIC = "public"
var ROOM_TYPE_PRIVATE = "private"

type Room struct {
	ID       string
	Code     string
	Name     string
	Type     string
	Messages []vo.Message
}

func NewRoom(roomName, roomType string) (*Room, error) {
	code, err := shortid.Generate()
	if err != nil {
		return &Room{}, errors.Wrap(err, fmt.Sprintf("error on generating shortid for room[%s,%s]", roomName, roomType))
	}

	id := uuid.Must(uuid.NewV4(), nil).String()

	room := &Room{
		ID:       id,
		Name:     roomName,
		Code:     code,
		Type:     roomType,
		Messages: make([]vo.Message, 0),
	}

	return room, nil
}

func (r Room) Thumbnail() *vo.RoomThumbnail {
	var recentMsg []vo.Message
	if len(r.Messages) < 10 {
		recentMsg = r.Messages
	} else {
		recentMsg = r.Messages[(len(r.Messages) - 10):] // recent 10 messages
	}

	return &vo.RoomThumbnail{
		RoomID:         r.ID,
		RoomName:       r.Name,
		RoomType:       r.Type,
		RecentMessages: recentMsg,
	}
}
