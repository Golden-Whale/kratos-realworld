package biz

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
	"kratos-realworld/internal/conf"
	"kratos-realworld/pkg/middleware/auth"
)

type User struct {
	Email        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

type UserLogin struct {
	Email    string
	Token    string
	Username string
	Bio      string
	Image    string
}

func hashPassword(pwd string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", b)
	return string(b)
}

func verifyPassword(hashed, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
	if err != nil {
		return false
	}
	return true
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type ProfileRepo interface {
}

type UserUsecase struct {
	ur UserRepo
	pr ProfileRepo

	jwtConf *conf.JWT
	log     *log.Helper
}

func NewUserUsecase(ur UserRepo, pr ProfileRepo, jwtConf *conf.JWT, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, pr: pr, jwtConf: jwtConf, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) generateToken(username string) string {
	return auth.GenerateToken(uc.jwtConf.Secret, username)
}

func (uc *UserUsecase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	user := &User{
		Email:        email,
		Username:     username,
		PasswordHash: hashPassword(password),
	}
	err := uc.ur.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &UserLogin{
		Email:    email,
		Token:    uc.generateToken(user.Username),
		Username: username,
	}, nil
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*UserLogin, error) {
	u, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if verifyPassword(u.PasswordHash, password) {
		return nil, errors.New("login failed")
	}
	return &UserLogin{
		Email:    u.Email,
		Token:    uc.generateToken(u.Username),
		Username: u.Username,
		Bio:      u.Bio,
		Image:    u.Image,
	}, nil
}
