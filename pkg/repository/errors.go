package repository

import "errors"

var ErrorUsernameAlreadyTaken = errors.New("username is already taken")
var ErrorNoRows = errors.New("no rows found for this query")
