package repo

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uint
	Name     string
	Email    string
	Password []byte
	Telegram int64
}

func (user *User) CreatePassword(password string) {

	user.Password, _ = bcrypt.GenerateFromPassword([]byte(password), 14)

}

func (user *User) CheckPassword(password string) error {

	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
