package usecase

import "errors"

var ErrFieldsValidation = errors.New("failure fields validation")
var ErrIDValidation = errors.New("failure id validation")
var ErrNoExistsData = errors.New("no exists data")

// user
var ErrUserNameAlreadyUsed = errors.New("user name already used")
