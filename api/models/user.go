package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/matkinhig/go-blogs/api/security"
)

type User struct {
	ID        uint32    `gorm:"primary_key; auto_increment" json:"id"`
	Nickname  string    `gorm:"size:20; not null; unique" json:"nickname"`
	Email     string    `gorm:"size:50; not null; unique" json:"email"`
	Password  string    `gorm:"size:100; not null; unique" json:"password"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	Posts     []Post    `gorm:"foreignkey:AuthorID" json:"posts,omitempty"`
}

func (u *User) BeforeSave() error {
	hp, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hp)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(a string) error {
	switch strings.ToLower(a) {
	case "update":
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
		return nil
	default:
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
		return nil
	}

}
