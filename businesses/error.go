package businesses

import "errors"

var (
	ErrInternalServer           = errors.New("Something Gone Wrong,Please Contact Administrator")
	ErrPassword                 = errors.New("Wrong Password")
	ErrDuplicateUsername        = errors.New("Username already used")
	ErrUsernameNotRegistered    = errors.New("Username Not Registered")
	ErrUsernamePasswordNotFound = errors.New("(Email) Or (Password) Empty")
	ErrNotFoundTransaction      = errors.New("Transaction Does'nt Exist")
)
