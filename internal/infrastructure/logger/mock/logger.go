package mock

import "github.com/stretchr/testify/mock"

// MockLogger для тестов
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Debug(msg string, fields ...any) {
	if len(fields) == 0 {
		m.Called(msg)
	} else {
		m.Called(msg, fields)
	}
}

func (m *MockLogger) Info(msg string, fields ...any) {
	if len(fields) == 0 {
		m.Called(msg)
	} else {
		m.Called(msg, fields)
	}
}

func (m *MockLogger) Warn(msg string, fields ...any) {
	if len(fields) == 0 {
		m.Called(msg)
	} else {
		m.Called(msg, fields)
	}
}

func (m *MockLogger) Error(msg string, fields ...any) {
	if len(fields) == 0 {
		m.Called(msg)
	} else {
		m.Called(msg, fields)
	}
}

func (m *MockLogger) Fatal(msg string, fields ...any) {
	if len(fields) == 0 {
		m.Called(msg)
	} else {
		m.Called(msg, fields)
	}
}
