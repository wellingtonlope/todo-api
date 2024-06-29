package clock

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type ClientMock struct {
	mock.Mock
}

func NewClientMock() *ClientMock {
	return new(ClientMock)
}

func (m *ClientMock) Now() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}
