package model

import (
	"database/sql"
	"time"
)

type EnglishAnswer struct {
	ID        int64         `bun:"id,autoincrement,notnull"` // SERIAL primary key
	UserID    string        `bun:"user_id,type:varchar(255),notnull"`
	Sentence  string        `bun:"sentence,type:varchar(255),notnull"`
	ParentID  sql.NullInt64 `bun:"parent_id,on_delete:cascade"`
	CreatedAt time.Time     `bun:"created_at,default:current_timestamp"`
	UpdatedAt time.Time     `bun:"updated_at,default:current_timestamp"`
}
