package exceptions

import "errors"

// type MinLengthInvalidPasswordError struct {
// 	Title  string
// 	Detail string
// 	// Err    error
// }

// func (e *MinLengthInvalidPasswordError) Error() string {
// 	return "enter_a_valid_password1"
// }

var UserLoginFailedError = errors.New("user_login_failed")
