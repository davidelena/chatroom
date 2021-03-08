package model

import "errors"

// Define some custom error code according to the business

var (
	ERROR_INTERNAL_SERVER_ERROR = errors.New("server internal error")
	ERROR_USER_NOT_EXIST        = errors.New("user is not existed")
	ERROR_USER_EXISTS           = errors.New("user is exsited")
	ERROR_USER_PASSWORD         = errors.New("user password is not corrected")
	ERROR_USER_NOT_ONLINE       = errors.New("user is not online")
)
