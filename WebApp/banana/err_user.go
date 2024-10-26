package banana

import "errors"

var (
	UserConflict= errors.New("user already exits!")
)