package database

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:user"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Name          string `bun:"name"`
	Username      string `bun:"username"`
	Email         string `bun:"email"`
	PasswordHash  string `bun:"password_hash"`
	AuthProvider  string `bun:"auth_provider"`
	AuthID        string `bun:"auth_id"`
	Avatar        string `bun:"avatar"`
	IsAdmin       bool   `bun:"is_admin"`
}

type Session struct {
	bun.BaseModel `bun:"table:session"`
	ID            string    `bun:"id,pk"`
	ExpiresAt     time.Time `bun:"expires_at"`
	UserID        int       `bun:"user_id"`
	User          *User     `bun:"rel:belongs-to,join:user_id=id"`
}

type Organization struct {
	bun.BaseModel `bun:"table:organization"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Name          string `bun:"name"`
}

type OrganizationMember struct {
	bun.BaseModel  `bun:"table:organization_member"`
	ID             int64         `bun:"id,pk,autoincrement"`
	Role           string        `bun:"role"`
	UserID         int64         `bun:"user_id"`
	OrganizationID int64         `bun:"organization_id"`
	User           *User         `bun:"rel:belongs-to,join:user_id=id"`
	Organization   *Organization `bun:"rel:belongs-to,join:organization_id=id"`
}

type Token struct {
	bun.BaseModel `bun:"table:token"`
	ID            int64  `bun:"id,pk,autoincrement"`
	TokenHash     string `bun:"token_hash"`
	Identifier    string `bun:"identifier"`
	UserID        int    `bun:"user_id"`
	User          *User  `bun:"rel:belongs-to,join:user_id=id"`
}

type File struct {
	bun.BaseModel  `bun:"table:file"`
	Key            string        `bun:"key"`
	Size           int64         `bun:"size"`
	Type           string        `bun:"type"`
	OrganizationID int64         `bun:"organization_id"`
	Organization   *Organization `bun:"rel:belongs-to,join:organization_id=id"`
}

type Store interface {
	Connect(dsn string) *bun.DB
}

// TODO: Return error if driver is not supported
func New(driver string) Store {
	switch driver {
	case "mysql":
		return &MySQL{}
	case "pg", "postgresql":
		return &PostgreSQL{}
	default:
		return nil
	}
}
