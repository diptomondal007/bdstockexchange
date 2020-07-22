package bdstockexchange

import "errors"

var (
	// ErrInvalidGroupName is thrown when user input an illegal group name
	ErrInvalidGroupName = errors.New("group name is invalid. enter a valid group name. ex : A, B, G, N, Z")
)
