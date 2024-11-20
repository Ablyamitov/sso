package auth

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/ablyamiov/sso/internal/domain/models"
	"github.com/ablyamiov/sso/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Auth struct {
	userSaver    UserSaver
	userProvider UserProvider
	appProvider  AppProvider
	tokenTTL     time.Duration
}

type UserSaver interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (userId int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userId int) (bool, error)
}

type AppProvider interface {
	App(ctx context.Context, appId int) (models.App, error)
}

func New(
	userSaver UserSaver,
	userProvider UserProvider,
	appProvider AppProvider,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		userSaver:    userSaver,
		userProvider: userProvider,
		appProvider:  appProvider,
		tokenTTL:     tokenTTL,
	}
}

func (a *Auth) Login(ctx context.Context, email string, password string, appId int) (string, error) {
	log.Println("login user")
	user, err := a.userProvider.User(ctx, email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			log.Println("user not found")
			return "", ErrInvalidCredentials
		}
		log.Println("failed to get user")
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		log.Printf("invalid credentials: %s", err)
		return "", ErrInvalidCredentials
	}
	app, err := a.appProvider.App(ctx, appId)
	if err != nil {
		log.Println("failed to get app")
		return "", err
	}

	//TODO: implement
	return "", nil
}

func (a *Auth) RegisterNewUser(ctx context.Context, email string, password string) (int64, error) {
	log.Println("register new user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("failed to hash password")
		return 0, err
	}

	id, err := a.userSaver.SaveUser(ctx, email, passHash)
	if err != nil {
		log.Println("failed to save user")
		return 0, err
	}
	return id, nil

}

func (a *Auth) IsAdmin(ctx context.Context, userId int) (bool, error) {
	//TODO: implement
	return false, nil
}
