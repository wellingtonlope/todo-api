package clock

import "time"

type Client interface {
	Now() time.Time
}
