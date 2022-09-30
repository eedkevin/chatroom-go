package room

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type CreateRoomArgs struct {
	Name string
}

type BroadcastMessageArgs struct {
	From    string
	To      string
	Content string
}

func (r *CreateRoomArgs) LoadFromJSON(data []byte) error {
	err := json.Unmarshal(data, &r)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error on CreateUserArgs.LoadFromJSON, %v", string(data)))
	}
	return nil
}

func (r *BroadcastMessageArgs) LoadFromJSON(data []byte) error {
	err := json.Unmarshal(data, &r)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error on BroadcastMessageArgs.LoadFromJSON, %v", string(data)))
	}
	return nil
}
