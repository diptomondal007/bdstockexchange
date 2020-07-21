package bdstockexchange

import "errors"

var (
	ErrInvalidGroupName = errors.New("group name is invalid. enter a valid group name. ex : A, B, G, N, Z")
)
