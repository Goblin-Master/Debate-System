package model

type User struct {
	ID       int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	UserID   int64  `gorm:"column:user_id;unique;type:bigint;not null"`
	Account  string `gorm:"column:account;unique;type:varchar(64);not null"`
	Nickname string `gorm:"column:nickname;type:varchar(32);not null"`
	Password string `gorm:"column:password;type:varchar(64);not null"`
	Avatar   string `gorm:"column:avatar;type:varchar(256);not null"`
	Utime    int64  `gorm:"column:utime;type:bigint;not null"`
	Ctime    int64  `gorm:"column:ctime;type:bigint;not null"`
}

func (User) TableName() string {
	return "user"
}
