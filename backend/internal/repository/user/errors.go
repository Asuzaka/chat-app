package user

import "errors"

var NoUserFoundError = errors.New("no user found")
var DuplicateUserError = errors.New("duplicate user")
var InvalidCredentialsError = errors.New("invalid credentials")
