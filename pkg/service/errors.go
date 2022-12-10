package service

import "errors"

var ErrorInvalidCredentials = errors.New("invalid credentials")
var ErrorUserNotFound = errors.New("user not found")
