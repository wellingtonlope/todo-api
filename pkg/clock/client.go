package clock

import "time"

type (
	Client interface {
		Now() time.Time
	}
	client struct{}
)

func NewClient() *client {
	return &client{}
}

func (c *client) Now() time.Time {
	return time.Now().UTC()
}
