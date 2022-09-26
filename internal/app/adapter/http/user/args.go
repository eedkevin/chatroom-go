package user

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type CreateUserArgs struct {
	Name string
}

func (u *CreateUserArgs) LoadFromJSON(data []byte) error {
	err := json.Unmarshal(data, &u)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error on CreateUserArgs.LoadFromJSON, %v", string(data)))
	}
	return nil
}
