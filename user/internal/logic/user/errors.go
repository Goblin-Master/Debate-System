package user

import "errors"

var (
	USER_NOT_EXIST       = errors.New("用户不存在")
	ACCOUNT_OR_PWD_ERROR = errors.New("账号或密码错误")
	ACCOUNT_EXIST        = errors.New("账号已经存在")
	DEFAULT_ERROR        = errors.New("默认错误")
	VAILD_EMPTY          = errors.New("参数为空")
	MODIFY_ERROR         = errors.New("修改失败")
)
