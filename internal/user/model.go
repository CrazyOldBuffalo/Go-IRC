package user

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	username string
	bio      string
	age      int16
	userid   uuid.UUID
}

func NewUser(username string, bio string, age int16) (*User, error) {
	if err := validate(username, bio, age); err != nil {
		return nil, err
	}
	return &User{
		username: username,
		bio:      bio,
		age:      age,
		userid:   uuid.New(),
	}, nil
}

func validate(username string, bio string, age int16) error {
	if len(username) <= 2 || len(username) >= 20 {
		return fmt.Errorf("invalid username length: %d, is smaller than 2 or greater than 20", len(username))
	}
	if age <= 16 || age >= 120 {
		return fmt.Errorf("age: %d, is smaller than 16, or greater than 120", age)
	}
	return nil
}
