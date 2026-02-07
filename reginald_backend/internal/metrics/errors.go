package metrics

import "errors"

var (
	ErrUnauthorized = errors.New("only admins can view metrics")
)
