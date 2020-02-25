package models

import (
	"time"

	"github.com/matkinhig/go-blogs/api/security"
)

type User struct {
	ID        uint32    `gorm:"primary_key; auto_increment" json:"id"`
	Nickname  string    `gorm:"size:20; not null; unique" json:"nickname"`
	Email     string    `gorm:"size:50; not null; unique" json:"email"`
	Password  string    `gorm:"size:100; not null; unique" json:"password"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

func (u *User) BeforeSave() error {
	hp, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hp)
	return nil
}
