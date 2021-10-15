// Matheus Leonel Balduino
// Everywhere, under @mathbalduino
//   @mathbalduino on GitHub
//   @mathbalduino on Instagram
//   @mathbalduino on Twitter
// Live at mathbalduino.com.br
// 2021-10-15 1:21 PM

package parser

type mockLoggerCLI struct {
	mockTrace     func(format string, args ...interface{}) LoggerCLI
	mockDebug     func(format string, args ...interface{}) LoggerCLI
	mockInfo      func(format string, args ...interface{}) LoggerCLI
	mockWarn      func(format string, args ...interface{}) LoggerCLI
	mockError     func(format string, args ...interface{}) LoggerCLI
	mockFatal     func(format string, args ...interface{})
	mockErrorFrom func(e error) LoggerCLI
	mockFatalFrom func(e error)
}

func (m *mockLoggerCLI) Trace(format string, args ...interface{}) LoggerCLI {
	if m.mockTrace != nil {
		m.mockTrace(format, args...)
	}
	return m
}
func (m *mockLoggerCLI) Debug(format string, args ...interface{}) LoggerCLI {
	if m.mockDebug != nil {
		m.mockDebug(format, args...)
	}
	return m
}
func (m *mockLoggerCLI) Info(format string, args ...interface{}) LoggerCLI {
	if m.mockInfo != nil {
		m.mockInfo(format, args...)
	}
	return m
}
func (m *mockLoggerCLI) Warn(format string, args ...interface{}) LoggerCLI {
	if m.mockWarn != nil {
		m.mockWarn(format, args...)
	}
	return m
}
func (m *mockLoggerCLI) Error(format string, args ...interface{}) LoggerCLI {
	if m.mockError != nil {
		m.mockError(format, args...)
	}
	return m
}
func (m *mockLoggerCLI) Fatal(format string, args ...interface{}) {
	if m.mockFatal != nil {
		m.mockFatal(format, args...)
	}
}
func (m *mockLoggerCLI) ErrorFrom(e error) LoggerCLI {
	if m.mockErrorFrom != nil {
		m.mockErrorFrom(e)
	}
	return m
}
func (m *mockLoggerCLI) FatalFrom(e error) {
	if m.mockFatalFrom != nil {
		m.mockFatalFrom(e)
	}
}
