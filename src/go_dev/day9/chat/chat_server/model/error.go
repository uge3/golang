package model

import "errors"

//主定议错误
var (
	ErrUserNotExist  = errors.New("user not exist")
	ErrInvalidPasswd = errors.New("passwd or username not right")
	ErrInvalidParams = errors.New("Invalid paramas")
	ErrUserExist     = errors.New("user exist")
)
