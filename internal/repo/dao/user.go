package dao

import (
	"Debate-System/internal/model"
	"Debate-System/internal/svc"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

type UserDao struct {
	svcCtx *svc.ServiceContext
	ctx    context.Context
}

var (
	ACCOUNT_EXIST = errors.New("账号已经存在")
)

func NewUserDao(svcCtx *svc.ServiceContext, ctx context.Context) *UserDao {
	return &UserDao{
		svcCtx: svcCtx,
		ctx:    ctx,
	}
}
func (ud *UserDao) GetByID(user_id int64) (model.User, error) {
	var user model.User
	err := ud.svcCtx.DB.WithContext(ud.ctx).Where("user_id = ?", user_id).Take(&user).Error
	return user, err
}

func (ud *UserDao) GetByAccount(account string) (model.User, error) {
	var user model.User
	err := ud.svcCtx.DB.WithContext(ud.ctx).Where("account = ?", account).Take(&user).Error
	return user, err
}

func (ud *UserDao) Insert(account, pwd, name, avatar string, id int64) error {
	_, err := ud.GetByAccount(account)
	switch err {
	case nil:
		return ACCOUNT_EXIST
	case gorm.ErrRecordNotFound:
		t := time.Now().UnixMilli()
		err = ud.svcCtx.DB.WithContext(ud.ctx).Create(&model.User{
			Account:  account,
			Password: pwd,
			Nickname: name,
			UserID:   id,
			Avatar:   avatar,
			Ctime:    t,
			Utime:    t,
		}).Error
		return err
	default:
		return err
	}
}

func (ud *UserDao) CheckAccountAndPwd(account, pwd string) (model.User, error) {
	var user model.User
	err := ud.svcCtx.DB.WithContext(ud.ctx).Where("account = ? and password = ?", account, pwd).Take(&user).Error
	return user, err
}

func (ud *UserDao) UpdateData(user_id int64, name, avatar string) error {
	err := ud.svcCtx.DB.WithContext(ud.ctx).Model(&model.User{}).Where("user_id = ?", user_id).Updates(model.User{
		Nickname: name,
		Avatar:   avatar,
		Utime:    time.Now().UnixMilli(),
	}).Error
	return err
}
