package auth

import "time"

const (
	SessionCookieName = "fmlite_session"
	UserContextKey    = "user"
	OrgIDContextKey   = "org_id"
	SessionDuration   = 24 * time.Hour
)
