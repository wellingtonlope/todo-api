package usecase

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type ClockMock struct {
	mock.Mock
}

func NewClockMock() *ClockMock {
	return new(ClockMock)
}

func (m *ClockMock) Now() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}
