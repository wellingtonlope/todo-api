package clock

import "time"

type ClientUTC struct{}

func NewClientUTC() *ClientUTC {
	return &ClientUTC{}
}

func (c *ClientUTC) Now() time.Time {
	return time.Now().UTC()
}
