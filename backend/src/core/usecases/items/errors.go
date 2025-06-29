package items

import "errors"

var ErrItemIDRequired = errors.New("item ID is required")
var ErrItemNotFound = errors.New("item not found")
