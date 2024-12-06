package authservice

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/auth"
	"github.com/fivemanage/lite/internal/crypt"
	"github.com/fivemanage/lite/internal/database"
	"github.com/uptrace/bun"
	"golang.org/x/oauth2"
)

type authConfig struct {
	github *oauth2.Config
}

type Auth struct {
	config authConfig
	db     *bun.DB
}

func New(db *bun.DB) *Auth {
	githubConfig := auth.NewGithubConfig()

	return &Auth{
		config: authConfig{
			github: githubConfig,
		},
		db: db,
	}
}

// RegisterUser will register the user, not shit, but we need to make sure this is
// only for non-admin users. The actual admin user should have the defalt admin password
// unless changed in ENV, and then be prompted to change it. That should probably be required.
func (a *Auth) RegisterUser(ctx context.Context, register *api.RegisterRequest) (string, error) {
	exists := a.userExists(ctx, register.Email)
	if exists {
		return "", errors.New("user already exists")
	}

	userID, err := a.createUser(ctx, register)
	if err != nil {
		return "", err
	}

	sessionID, err := a.createSession(ctx, userID)
	if err != nil {
		return "", err
	}

	// not sure if we really need to return anything else than the error here
	// we can instead create a session as another call and then add that sessionId
	// as a cookie
	return sessionID, nil
}

// Login uses email and password to authenticate the user
func (a *Auth) LoginUser() {
}

// OAuthLogin uses OAuth2 to authenticate the user
// This will probably work as register and login, right???
func (a *Auth) OAuthLogin() string {
	verifier := oauth2.GenerateVerifier()
	url := a.config.github.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))

	return url
}

func (a *Auth) Callback(code string) *oauth2.Token {
	token, err := a.config.github.Exchange(context.TODO(), code)
	if err != nil {
		fmt.Println(err)
	}

	return token
}

func (a *Auth) userExists(ctx context.Context, email string) bool {
	user := new(database.User)
	err := a.db.NewSelect().Model(user).Where("email = ?", email).Scan(ctx)
	if err != nil {
		fmt.Println(err)
	}

	if user.ID == 0 {
		return false
	}

	return user.ID != 0
}

// CreateUser creates a new user with email and password
func (a *Auth) createUser(ctx context.Context, register *api.RegisterRequest) (int64, error) {
	var err error
	hash, err := crypt.HashPassword(register.Password)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	user := &database.User{
		Email:        register.Email,
		PasswordHash: hash,
	}

	// this should be a tx, so we can rollback if we fail to get the LID
	_, err = a.db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return 0, err
	}

	return 3, nil

	// return userID, nil
}

// not sure if this should be public or not yet
func (a *Auth) createSession(ctx context.Context, userID int64) (string, error) {
	sessionID, err := crypt.GenerateSessionID()
	if err != nil {
		return "", err
	}

	fmt.Println("do we get session id", sessionID)

	session := &database.Session{
		ID:        sessionID,
		UserID:    int(userID),
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	_, err = a.db.NewInsert().Model(session).Exec(ctx)
	if err != nil {
		fmt.Println("oh my god", err)
		return "", nil
	}

	fmt.Println("is the session id here", session.ID)

	return session.ID, nil
}
