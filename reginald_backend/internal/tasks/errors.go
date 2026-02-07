package tasks

import "errors"

var ErrTaskNotFound = errors.New("task not found")
var ErrInvalidStatus = errors.New("invalid status")
