package global

import (
	snowflake "Debate-System/utils/snowfake"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

var Node, _ = snowflake.NewNode(1)
