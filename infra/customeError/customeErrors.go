package customeError

import (
	"errors"
)

var UserNotFound error
var UserAlreadyExists error
var WrongCredentials error
var SomethingWentWrong error
var TodoNotFound error
var TodoCouldntCreate error
var TodoCouldntDelete error
var TodoCouldntUpdate error

func init() {
	UserNotFound = errors.New("UserNotFound-Err")
	UserAlreadyExists = errors.New("UserAlreadyExists-Err")
	WrongCredentials = errors.New("WrongCredentials-Err")
	SomethingWentWrong = errors.New("SomethingWentWrong-Err")
	TodoNotFound = errors.New("TodoNotFound-Err")
	TodoCouldntCreate = errors.New("TodoCouldntCreate-Err")
	TodoCouldntDelete = errors.New("TodoCouldntDelete-Err")
	TodoCouldntUpdate = errors.New("TodoCouldntUpdate-Err")
}
