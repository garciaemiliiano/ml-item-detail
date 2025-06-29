package rest

import "time"

type HandlerBase struct {
	CustomNow func() time.Time
}

func (h HandlerBase) Now() time.Time {
	if h.CustomNow != nil {
		return h.CustomNow()
	}
	return time.Now()
}
