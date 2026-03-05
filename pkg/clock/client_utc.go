package clock

import "time"

// ClientUTC provides time operations in UTC timezone.
type ClientUTC struct{}

// NewClientUTC returns a new ClientUTC instance.
func NewClientUTC() *ClientUTC {
	return &ClientUTC{}
}

// Now returns the current time in UTC.
//
// Returns:
//   - time.Time: The current UTC time.
func (c *ClientUTC) Now() time.Time {
	return time.Now().UTC()
}
