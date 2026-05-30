package tracker

import "errors"

var ErrNotFound = errors.New("not found")
var ErrIdNoUnique = errors.New("id is not unique")
