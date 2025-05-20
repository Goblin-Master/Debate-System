package repo

import (
	"Debate-System/internal/model"
	"Debate-System/internal/repo/dao"
	"Debate-System/internal/svc"
	"Debate-System/utils/snowfake"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type IUserRepo interface {
	GetUserByID(user_id int64) (model.User, error)
	CreateUser(account, pwd, name, avatar string) (int64, error)
	CheckLogin(account, pwd string) (model.User, error)
	ModifyUserData(user_id int64, name, avatar string) error
}
type UserRepo struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	userDao *dao.UserDao
}

var _ IUserRepo = (*UserRepo)(nil)

var (
	USER_NOT_EXIST       = errors.New("用户不存在")
	ACCOUNT_EXIST        = errors.New("账号已经存在")
	DEFAULT_ERROR        = errors.New("默认错误")
	ACCOUNT_OR_PWD_ERROR = errors.New("账号或密码错误")
)

func NewUserRepo(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepo {
	return &UserRepo{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		userDao: dao.NewUserDao(svcCtx),
	}
}

func (u *UserRepo) GetUserByID(user_id int64) (model.User, error) {
	user, err := u.userDao.GetByID(user_id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, USER_NOT_EXIST
		}
		u.Logger.Error(err)
		return model.User{}, DEFAULT_ERROR
	}
	return user, nil
}

func (u *UserRepo) CreateUser(account, pwd, name, avatar string) (int64, error) {
	id := snowfake.GetIntId(u.svcCtx.Node)
	err := u.userDao.Insert(account, pwd, name, avatar, id)
	if err != nil {
		if errors.Is(err, dao.ACCOUNT_EXIST) {
			return 0, ACCOUNT_EXIST
		}
		u.Logger.Error(err)
		return 0, DEFAULT_ERROR
	}
	return id, nil
}
func (u *UserRepo) CheckLogin(account, pwd string) (model.User, error) {
	user, err := u.userDao.CheckAccountAndPwd(account, pwd)
	if err != nil {
		return model.User{}, ACCOUNT_OR_PWD_ERROR
	}
	return user, nil
}
func (u *UserRepo) ModifyUserData(user_id int64, name, avatar string) error {
	_, err := u.userDao.GetByID(user_id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return USER_NOT_EXIST
		}
		u.Logger.Error(err)
		return DEFAULT_ERROR
	}
	err = u.userDao.UpdateData(user_id, name, avatar)
	if err != nil {
		u.Logger.Error(err)
		return DEFAULT_ERROR
	}
	return nil
}
