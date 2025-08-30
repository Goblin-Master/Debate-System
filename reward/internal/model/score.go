package model

type Score struct {
	ID     int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	UserID int64 `gorm:"column:user_id;unique;type:bigint;not null"`
	Utime  int64 `gorm:"column:utime;type:bigint;not null"`
	Ctime  int64 `gorm:"column:ctime;type:bigint;not null"`
}

func (Score) TableName() string {
	return "score"
}
