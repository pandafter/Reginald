package requests

import "errors"

var (
	ErrRequestNotFound = errors.New("request not found")
	ErrInvalidStatus   = errors.New("invalid request status: must be pending, review, approved or rejected")
	ErrUnauthorized    = errors.New("unauthorized: operation not allowed for this role")
	ErrInvalidBody     = errors.New("invalid request body")
)
