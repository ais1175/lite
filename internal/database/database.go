package database

import (
	"fmt"
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
	ID            string `bun:"id,pk"`
	OwnerID       string `bun:"owner_Id"`
	Name          string `bun:"name"`
	User          *User  `bun:"rel:belongs-to,join:owner_id=id"`
}

type OrganizationMember struct {
	bun.BaseModel  `bun:"table:organization_member"`
	ID             int64         `bun:"id,pk,autoincrement"`
	UserID         int           `bun:"user_id"`
	OrganizationID int           `bun:"organization_id"`
	User           *User         `bun:"rel:belongs-to,join:user_id=id"`
	Organization   *Organization `bun:"rel:belongs-to,join:organization_id=id"`
}

type Store interface {
	Connect() *bun.DB
}

// TODO: Return error if driver is not supported
func New(driver string, dsn string) Store {
	fmt.Println("option", driver)

	switch driver {
	case "mysql":
		return &MySQL{}
	case "sqlite":
		return &SQLite{}
	default:
		return nil
	}
}
