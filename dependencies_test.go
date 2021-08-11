package goParser

func emptyMockLogCLI() LogCLI {
	m := &mockLogCLI{}
	m.mockDebug = func(msgFormat string, args ...interface{}) LogCLI { return m }
	m.mockError = func(msgFormat string, args ...interface{}) LogCLI { return m }
	return m
}

type mockLogCLI struct {
	mockDebug func(msgFormat string, args ...interface{}) LogCLI
	mockError func(msgFormat string, args ...interface{}) LogCLI
}

func (m *mockLogCLI) Debug(msgFormat string, args ...interface{}) LogCLI {
	if m.mockDebug == nil {
		return nil
	}
	return m.mockDebug(msgFormat, args...)
}

func (m *mockLogCLI) Error(msgFormat string, args ...interface{}) LogCLI {
	if m.mockError == nil {
		return nil
	}
	return m.mockError(msgFormat, args...)
}
