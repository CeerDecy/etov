package apierrors

import "errors"

var (
	UserNotExistError         = errors.New("当前用户不存在，请先注册")
	UserPasswordNotMatchError = errors.New("密码错误，请重新输入")
)
