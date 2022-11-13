package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

// BeforeCreate ...
func (u *User) BeforeCreate() error {
	if len(u.EncryptedPassword) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return nil
		}
		u.EncryptedPassword = enc
	}
	return nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
