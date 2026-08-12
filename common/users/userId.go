package users

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type UserId uuid.UUID

func (u UserId) String() string {
	return uuid.UUID(u).String()
}

func (u UserId) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

func (u *UserId) UnmarshalJSON(data []byte) error {
	uid, err := uuid.Parse(string(data))
	if err != nil {
		return err
	}
	*u = UserId(uid)
	return nil
}

func (u *UserId) ToString() (string, error) {
	return u.String(), nil
}

func (u *UserId) FromString(s string) error {
	uid, err := uuid.Parse(s)
	if err != nil {
		return err
	}
	*u = UserId(uid)
	return nil
}

func (u *UserId) Scan(value any) error {
	if value == nil {
		*u = UserId{}
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("UserId: expected string, got %T", value)
	}
	u.FromString(str)
	return nil
}

func (u UserId) Value() (driver.Value, error) {
	return u.String(), nil
}
