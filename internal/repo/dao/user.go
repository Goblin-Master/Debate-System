package dao

import (
	"Debate-System/internal/global"
	"Debate-System/internal/model"
	"errors"
	"gorm.io/gorm"
	"sync"
)

type UserDao struct {
	db *gorm.DB
}

var userDao *UserDao
var once sync.Once

var (
	ACCOUNT_EXIST = errors.New("账号已经存在")
)

func NewUserDao() *UserDao {
	once.Do(func() {
		userDao = &UserDao{
			db: global.DB,
		}
	})
	return userDao
}
func (ud *UserDao) GetByID(id int64) (model.User, error) {
	var user model.User
	err := ud.db.Where("id = ?", id).Take(&user).Error
	return user, err
}

func (ud *UserDao) GetByAccount(account string) (model.User, error) {
	var user model.User
	err := ud.db.Where("account = ?", account).Take(&user).Error
	return user, err
}

func (ud *UserDao) Insert(account, pwd, name string, id int64) error {
	_, err := ud.GetByAccount(account)
	switch err {
	case nil:
		return ACCOUNT_EXIST
	case gorm.ErrRecordNotFound:
		err = ud.db.Create(&model.User{
			Account:  account,
			Password: pwd,
			Nickname: name,
			UserID:   id,
		}).Error
		return err
	default:
		return err
	}
}
