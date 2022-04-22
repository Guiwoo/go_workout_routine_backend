package model

import "time"

type User_Type struct {
	ID        int64 `xorm:"id pk autoincr"`
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
