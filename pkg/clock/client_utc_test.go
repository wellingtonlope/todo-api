package clock_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wellingtonlope/todo-api/pkg/clock"
)

func TestClientUTC_Now(t *testing.T) {
	testCases := []struct {
		name   string
		result time.Time
	}{
		{
			name:   "should get time",
			result: time.Now().UTC(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := clock.NewClientUTC()
			result := c.Now()
			assert.Equal(t, tc.result.Format(time.DateTime), result.Format(time.DateTime))
		})
	}
}
