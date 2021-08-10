package goParser

// LogCLI is used to log the parser actions, using only
// two levels: Debug and Error
type LogCLI interface {
	Debug(msgFormat string, args ...interface{}) LogCLI
	Error(msgFormat string, args ...interface{}) LogCLI
}
