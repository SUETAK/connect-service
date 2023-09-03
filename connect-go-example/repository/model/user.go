package model

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID       string `bun:"id,pk,type:uuid"`
	Email    string `bun:"type:varchar(255)"`
	Password string `bun:"type:varchar(255)"`
}
