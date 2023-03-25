package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserId   int    `json:"user_id" query:"user_id"  db:"user_id" gorm:"primaryKey;autoIncrement:true"`
	Name     string `json:"name" query:"name"  db:"name"`
	Age      int    `json:"age" query:"age"  db:"age"`
	Email    string `json:"email" query:"email"  db:"email" gorm:"size:255;index:idx_email,unique"`
	Password string `json:"password,omitempty" query:"password" db:"password"`
	Address  string `json:"address" db:"address"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (u *User) VerifyPassword(hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(u.Password))
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return nil
}
