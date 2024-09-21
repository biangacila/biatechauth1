package valueobjects

import (
	"encoding/json"
	"github.com/google/uuid"
)

type UUID struct {
	Value uuid.UUID
}

func NewUUID() UUID {
	return UUID{Value: uuid.New()}
}

func FromString(id string) (UUID, error) {
	uuid1 := uuid.MustParse(id)
	return UUID{Value: uuid1}, nil
}

func (u UUID) String() string {
	return u.Value.String()
}

func (u UUID) GetValue() uuid.UUID {
	return u.Value
}

func (u UUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Value.String())
}

func (u *UUID) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	parsedUUID := uuid.MustParse(s)
	u.Value = parsedUUID
	return nil
}
