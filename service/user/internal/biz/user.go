package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

// 定义返回数据结构体

type User struct {
	ID          int64
	Mobile      string
	Password    string
	NickName    string
	Birthday    *time.Time
	Gender      string
	Role        int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	IsDeletedAt bool
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) Create(ctx context.Context, u *User) (*User, error) {
	return uc.repo.CreateUser(ctx, u)
}

//go:generate mockgen -destination=../mocks/mrepo/user.go -package=mrepo . UserRepo
type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
	ListUser(ctx context.Context, pageNum, pageSize int) ([]*User, int, error)
	UserByMobile(ctx context.Context, mobile string) (*User, error)
	GetUserById(ctx context.Context, id int64) (*User, error)
	UpdateUser(context.Context, *User) (bool, error)
	CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error)
}

//-----------------

func (uc *UserUseCase) List(ctx context.Context, pageNum, pageSize int) ([]*User, int, error) {
	return uc.repo.ListUser(ctx, pageNum, pageSize)
}

func (uc *UserUseCase) UserByMobile(ctx context.Context, mobile string) (*User, error) {
	return uc.repo.UserByMobile(ctx, mobile)
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, user *User) (bool, error) {
	return uc.repo.UpdateUser(ctx, user)
}

func (uc *UserUseCase) CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error) {
	return uc.repo.CheckPassword(ctx, password, encryptedPassword)
}

func (uc *UserUseCase) UserById(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetUserById(ctx, id)
}
