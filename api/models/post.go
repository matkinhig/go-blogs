package models

import (
	"errors"
	"html"
	"strings"
	"time"

)

type Post struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size :500;not null" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	Author    User      `gorm:"-" json:"author"`
	AuthorID  uint32    `gorm:"not null" json:"authorid"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

func (p *Post) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Post) Validate() error {
	if p.Title == "" {
		return errors.New("Required Title")
	}
	if p.Content == "" {
		return errors.New("Required Content")
	}
	if p.AuthorID < 1 {
		return errors.New("Required Author")
	}
	return nil
}
