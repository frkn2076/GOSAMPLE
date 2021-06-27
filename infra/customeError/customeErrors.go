package customeError

import (
	"errors"
)

var UserNotFound error
var UserAlreadyExists error
var WrongCredentials error
var SomethingWentWrong error

func init() {
	UserNotFound = errors.New("UserNotFound-Err")
	UserAlreadyExists = errors.New("UserAlreadyExists-Err")
	WrongCredentials = errors.New("WrongCredentials-Err")
	SomethingWentWrong = errors.New("SomethingWentWrong-Err")
}
