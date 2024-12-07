package error

import "errors"

var (
	ErrorNotFound error = errors.New("product not found")
)
