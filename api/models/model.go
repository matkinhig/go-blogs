package models

import "time"

//code next in here
type Model struct {
	ID         uint64
	Created_At time.Time
	Update_At  time.Time
}
