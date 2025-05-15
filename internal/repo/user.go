package repo

import (
	"Debate-System/internal/global"
	"Debate-System/internal/model"
	"Debate-System/internal/repo/dao"
	"Debate-System/utils/snowfake"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"sync"
)

type IUserRepo interface {
	GetUserByID(id int64) (model.User, error)
	CreateUser(account, pwd, name string) (int64, error)
}
type UserRepo struct {
	logx.Logger
	ctx     context.Context
	userDao *dao.UserDao
}

var _ IUserRepo = (*UserRepo)(nil)

var userRepo *UserRepo
var once sync.Once

var (
	USER_NOT_EXIST = errors.New("用户不存在")
	ACCOUNT_EXIST  = errors.New("账号已经存在")
	DEFAULT_ERROR  = errors.New("默认错误")
)

func NewUserRepo(ctx context.Context) *UserRepo {
	once.Do(func() {
		userRepo = &UserRepo{
			Logger:  logx.WithContext(ctx),
			ctx:     ctx,
			userDao: dao.NewUserDao(),
		}
	})
	return userRepo
}

func (u *UserRepo) GetUserByID(id int64) (model.User, error) {
	user, err := u.userDao.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, USER_NOT_EXIST
		}
		u.Logger.Error(err)
		return model.User{}, DEFAULT_ERROR
	}
	return user, nil
}

func (u *UserRepo) CreateUser(account, pwd, name string) (int64, error) {
	id := snowfake.GetIntId(global.Node)
	err := u.userDao.Insert(account, pwd, name, id)
	if err != nil {
		if errors.Is(err, dao.ACCOUNT_EXIST) {
			return 0, ACCOUNT_EXIST
		}
		u.Logger.Error(err)
		return 0, DEFAULT_ERROR
	}
	return id, nil
}
