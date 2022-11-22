package user

import "time"

type Entity struct {
	ID        int64
	Email     string
	Password  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Entity) TableName() string {
	return "users"
}
