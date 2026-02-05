package todo_test

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type clockMock struct {
	mock.Mock
}

func newClockMock() *clockMock {
	return new(clockMock)
}

func (m *clockMock) Now() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}
