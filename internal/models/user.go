package models

import (
	"time"

	"github.com/uptrace/bun"
)

// User represents the standard database model for our application
type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID        int64     `bun:"id,pk,autoincrement"`
	Name      string    `bun:"name,notnull"`
	LastName  string    `bun:"last_name,notnull"`
	Birthdate time.Time `bun:"birthdate"`
	Email     string    `bun:"email,notnull"`
	Password  string    `bun:"password,notnull"`
	IsActive  bool      `bun:"is_active,notnull,default:true"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}
