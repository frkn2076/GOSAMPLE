package customeError

var UserNotFound error
var UserAlreadyExists error
var WrongCredentials error
var RequestBodyIsNotValid error
var SomethingWentWrong error

func init() {
	UserNotFound = New(101, "UserNotFound-Err")
	UserAlreadyExists = New(102, "UserAlreadyExists-Err")
	WrongCredentials = New(103, "WrongCredentials-Err")

	RequestBodyIsNotValid = New(201, "SomethingWentWrong-Err")

	SomethingWentWrong = New(999, "SomethingWentWrong-Err")
}
