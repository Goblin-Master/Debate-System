package dao

import (
	"Debate-System/internal/model"
	"Debate-System/internal/svc"
	"errors"
	"gorm.io/gorm"
)

type UserDao struct {
	svcCtx *svc.ServiceContext
}

var (
	ACCOUNT_EXIST = errors.New("账号已经存在")
)

func NewUserDao(svcCtx *svc.ServiceContext) *UserDao {
	return &UserDao{
		svcCtx: svcCtx,
	}
}
func (ud *UserDao) GetByID(id int64) (model.User, error) {
	var user model.User
	err := ud.svcCtx.DB.Where("id = ?", id).Take(&user).Error
	return user, err
}

func (ud *UserDao) GetByAccount(account string) (model.User, error) {
	var user model.User
	err := ud.svcCtx.DB.Where("account = ?", account).Take(&user).Error
	return user, err
}

func (ud *UserDao) Insert(account, pwd, name, avatar string, id int64) error {
	_, err := ud.GetByAccount(account)
	switch err {
	case nil:
		return ACCOUNT_EXIST
	case gorm.ErrRecordNotFound:
		err = ud.svcCtx.DB.Create(&model.User{
			Account:  account,
			Password: pwd,
			Nickname: name,
			UserID:   id,
			Avatar:   avatar,
		}).Error
		return err
	default:
		return err
	}
}

func (ud *UserDao) CheckAccountAndPwd(account, pwd string) (model.User, error) {
	var user model.User
	err := ud.svcCtx.DB.Where("account = ? and password = ?", account, pwd).Take(&user).Error
	return user, err
}
