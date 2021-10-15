package parser

import (
	"fmt"
	"go/token"
	"reflect"
	"testing"
)

func TestPrintFinalConfig(t *testing.T) {
	t.Run("Print the correct final configuration when focus, fset and dir are nil", func(t *testing.T) {
		config := Config{
			true,
			"",
			[]string{},
			nil,
			[]string{},
			nil,
			0,
		}
		log := &mockLoggerCLI{}
		calls := 0
		pattern := "somePattern"
		log.mockDebug = func(msgFormat string, args ...interface{}) LoggerCLI {
			calls += 1

			if msgFormat != finalConfigTemplate {
				t.Fatalf("Wrong Log.Debug template")
			}
			if len(args) != 8 {
				t.Fatalf("Wrong Log.Debug variadic args")
			}
			if args[0].(string) != pattern {
				t.Fatalf("Wrong Log.Debug pattern")
			}
			if args[1].(bool) != config.Tests {
				t.Fatalf("Wrong Log.Debug tests boolean")
			}
			if args[2].(string) != emptyDirStr {
				t.Fatalf("Wrong Log.Debug dir string")
			}
			if !reflect.DeepEqual(args[3].([]string), config.Env) {
				t.Fatalf("Wrong Log.Debug Env slice")
			}
			if !reflect.DeepEqual(args[4].([]string), config.BuildFlags) {
				t.Fatalf("Wrong Log.Debug BuildFlags slice")
			}
			if args[5].(string) != nilFocusStr {
				t.Fatalf("Wrong Log.Debug focus string")
			}
			if args[6].(string) != nilFsetStr {
				t.Fatalf("Wrong Log.Debug Fset string")
			}
			if args[7].(string) != "-" {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration when focus and dir are nil", func(t *testing.T) {
		config := Config{
			true,
			"",
			[]string{},
			token.NewFileSet(),
			[]string{},
			nil,
			0,
		}
		log := &mockLoggerCLI{}
		calls := 0
		pattern := "somePattern"
		log.mockDebug = func(msgFormat string, args ...interface{}) LoggerCLI {
			calls += 1

			if msgFormat != finalConfigTemplate {
				t.Fatalf("Wrong Log.Debug template")
			}
			if len(args) != 8 {
				t.Fatalf("Wrong Log.Debug variadic args")
			}
			if args[0].(string) != pattern {
				t.Fatalf("Wrong Log.Debug pattern")
			}
			if args[1].(bool) != config.Tests {
				t.Fatalf("Wrong Log.Debug tests boolean")
			}
			if args[2].(string) != emptyDirStr {
				t.Fatalf("Wrong Log.Debug dir string")
			}
			if !reflect.DeepEqual(args[3].([]string), config.Env) {
				t.Fatalf("Wrong Log.Debug Env slice")
			}
			if !reflect.DeepEqual(args[4].([]string), config.BuildFlags) {
				t.Fatalf("Wrong Log.Debug BuildFlags slice")
			}
			if args[5].(string) != nilFocusStr {
				t.Fatalf("Wrong Log.Debug focus string")
			}
			if args[6].(string) != notNilFsetStr {
				t.Fatalf("Wrong Log.Debug Fset string")
			}
			if args[7].(string) != "-" {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration when focus is nil", func(t *testing.T) {
		config := Config{
			true,
			"abc",
			[]string{},
			token.NewFileSet(),
			[]string{},
			nil,
			0,
		}
		log := &mockLoggerCLI{}
		calls := 0
		pattern := "somePattern"
		log.mockDebug = func(msgFormat string, args ...interface{}) LoggerCLI {
			calls += 1

			if msgFormat != finalConfigTemplate {
				t.Fatalf("Wrong Log.Debug template")
			}
			if len(args) != 8 {
				t.Fatalf("Wrong Log.Debug variadic args")
			}
			if args[0].(string) != pattern {
				t.Fatalf("Wrong Log.Debug pattern")
			}
			if args[1].(bool) != config.Tests {
				t.Fatalf("Wrong Log.Debug tests boolean")
			}
			if args[2].(string) != config.Dir {
				t.Fatalf("Wrong Log.Debug dir string")
			}
			if !reflect.DeepEqual(args[3].([]string), config.Env) {
				t.Fatalf("Wrong Log.Debug Env slice")
			}
			if !reflect.DeepEqual(args[4].([]string), config.BuildFlags) {
				t.Fatalf("Wrong Log.Debug BuildFlags slice")
			}
			if args[5].(string) != nilFocusStr {
				t.Fatalf("Wrong Log.Debug focus string")
			}
			if args[6].(string) != notNilFsetStr {
				t.Fatalf("Wrong Log.Debug Fset string")
			}
			if args[7].(string) != "-" {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration", func(t *testing.T) {
		focusStr := "somePackagePath"
		config := Config{
			true,
			"abc",
			[]string{},
			token.NewFileSet(),
			[]string{},
			FocusPackagePath(focusStr),
			0,
		}
		log := &mockLoggerCLI{}
		calls := 0
		pattern := "somePattern"
		log.mockDebug = func(msgFormat string, args ...interface{}) LoggerCLI {
			calls += 1

			if msgFormat != finalConfigTemplate {
				t.Fatalf("Wrong Log.Debug template")
			}
			if len(args) != 8 {
				t.Fatalf("Wrong Log.Debug variadic args")
			}
			if args[0].(string) != pattern {
				t.Fatalf("Wrong Log.Debug pattern")
			}
			if args[1].(bool) != config.Tests {
				t.Fatalf("Wrong Log.Debug tests boolean")
			}
			if args[2].(string) != config.Dir {
				t.Fatalf("Wrong Log.Debug dir string")
			}
			if !reflect.DeepEqual(args[3].([]string), config.Env) {
				t.Fatalf("Wrong Log.Debug Env slice")
			}
			if !reflect.DeepEqual(args[4].([]string), config.BuildFlags) {
				t.Fatalf("Wrong Log.Debug BuildFlags slice")
			}
			if args[5].(string) != fmt.Sprintf(focusTemplate, focusStr, "nil", "nil", "nil", "nil") {
				t.Fatalf("Wrong Log.Debug focus string")
			}
			if args[6].(string) != notNilFsetStr {
				t.Fatalf("Wrong Log.Debug Fset string")
			}
			if args[7].(string) != "-" {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration", func(t *testing.T) {
		focusStr := "somePackagePath"
		config := Config{
			true,
			"abc",
			[]string{},
			token.NewFileSet(),
			[]string{},
			FocusFilePath(focusStr),
			0,
		}
		log := &mockLoggerCLI{}
		calls := 0
		pattern := "somePattern"
		log.mockDebug = func(msgFormat string, args ...interface{}) LoggerCLI {
			calls += 1

			if msgFormat != finalConfigTemplate {
				t.Fatalf("Wrong Log.Debug template")
			}
			if len(args) != 8 {
				t.Fatalf("Wrong Log.Debug variadic args")
			}
			if args[0].(string) != pattern {
				t.Fatalf("Wrong Log.Debug pattern")
			}
			if args[1].(bool) != config.Tests {
				t.Fatalf("Wrong Log.Debug tests boolean")
			}
			if args[2].(string) != config.Dir {
				t.Fatalf("Wrong Log.Debug dir string")
			}
			if !reflect.DeepEqual(args[3].([]string), config.Env) {
				t.Fatalf("Wrong Log.Debug Env slice")
			}
			if !reflect.DeepEqual(args[4].([]string), config.BuildFlags) {
				t.Fatalf("Wrong Log.Debug BuildFlags slice")
			}
			if args[5].(string) != fmt.Sprintf(focusTemplate, "nil", focusStr, "nil", "nil", "nil") {
				t.Fatalf("Wrong Log.Debug focus string")
			}
			if args[6].(string) != notNilFsetStr {
				t.Fatalf("Wrong Log.Debug Fset string")
			}
			if args[7].(string) != "-" {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration", func(t *testing.T) {
		focusStr := "somePackagePath"
		config := Config{
			true,
			"abc",
			[]string{},
			token.NewFileSet(),
			[]string{},
			FocusTypeName(focusStr),
			0,
		}
		log := &mockLoggerCLI{}
		calls := 0
		pattern := "somePattern"
		log.mockDebug = func(msgFormat string, args ...interface{}) LoggerCLI {
			calls += 1

			if msgFormat != finalConfigTemplate {
				t.Fatalf("Wrong Log.Debug template")
			}
			if len(args) != 8 {
				t.Fatalf("Wrong Log.Debug variadic args")
			}
			if args[0].(string) != pattern {
				t.Fatalf("Wrong Log.Debug pattern")
			}
			if args[1].(bool) != config.Tests {
				t.Fatalf("Wrong Log.Debug tests boolean")
			}
			if args[2].(string) != config.Dir {
				t.Fatalf("Wrong Log.Debug dir string")
			}
			if !reflect.DeepEqual(args[3].([]string), config.Env) {
				t.Fatalf("Wrong Log.Debug Env slice")
			}
			if !reflect.DeepEqual(args[4].([]string), config.BuildFlags) {
				t.Fatalf("Wrong Log.Debug BuildFlags slice")
			}
			if args[5].(string) != fmt.Sprintf(focusTemplate, "nil", "nil", focusStr, "nil", "nil") {
				t.Fatalf("Wrong Log.Debug focus string")
			}
			if args[6].(string) != notNilFsetStr {
				t.Fatalf("Wrong Log.Debug Fset string")
			}
			if args[7].(string) != "-" {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration", func(t *testing.T) {
		focusStr := "somePackagePath"
		config := Config{
			true,
			"abc",
			[]string{},
			token.NewFileSet(),
			[]string{},
			FocusVarName(focusStr),
			0,
		}
		log := &mockLoggerCLI{}
		calls := 0
		pattern := "somePattern"
		log.mockDebug = func(msgFormat string, args ...interface{}) LoggerCLI {
			calls += 1

			if msgFormat != finalConfigTemplate {
				t.Fatalf("Wrong Log.Debug template")
			}
			if len(args) != 8 {
				t.Fatalf("Wrong Log.Debug variadic args")
			}
			if args[0].(string) != pattern {
				t.Fatalf("Wrong Log.Debug pattern")
			}
			if args[1].(bool) != config.Tests {
				t.Fatalf("Wrong Log.Debug tests boolean")
			}
			if args[2].(string) != config.Dir {
				t.Fatalf("Wrong Log.Debug dir string")
			}
			if !reflect.DeepEqual(args[3].([]string), config.Env) {
				t.Fatalf("Wrong Log.Debug Env slice")
			}
			if !reflect.DeepEqual(args[4].([]string), config.BuildFlags) {
				t.Fatalf("Wrong Log.Debug BuildFlags slice")
			}
			if args[5].(string) != fmt.Sprintf(focusTemplate, "nil", "nil", "nil", focusStr, "nil") {
				t.Fatalf("Wrong Log.Debug focus string")
			}
			if args[6].(string) != notNilFsetStr {
				t.Fatalf("Wrong Log.Debug Fset string")
			}
			if args[7].(string) != "-" {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration", func(t *testing.T) {
		focusStr := "somePackagePath"
		config := Config{
			true,
			"abc",
			[]string{},
			token.NewFileSet(),
			[]string{},
			FocusFunctionName(focusStr),
			0,
		}
		log := &mockLoggerCLI{}
		calls := 0
		pattern := "somePattern"
		log.mockDebug = func(msgFormat string, args ...interface{}) LoggerCLI {
			calls += 1

			if msgFormat != finalConfigTemplate {
				t.Fatalf("Wrong Log.Debug template")
			}
			if len(args) != 8 {
				t.Fatalf("Wrong Log.Debug variadic args")
			}
			if args[0].(string) != pattern {
				t.Fatalf("Wrong Log.Debug pattern")
			}
			if args[1].(bool) != config.Tests {
				t.Fatalf("Wrong Log.Debug tests boolean")
			}
			if args[2].(string) != config.Dir {
				t.Fatalf("Wrong Log.Debug dir string")
			}
			if !reflect.DeepEqual(args[3].([]string), config.Env) {
				t.Fatalf("Wrong Log.Debug Env slice")
			}
			if !reflect.DeepEqual(args[4].([]string), config.BuildFlags) {
				t.Fatalf("Wrong Log.Debug BuildFlags slice")
			}
			if args[5].(string) != fmt.Sprintf(focusTemplate, "nil", "nil", "nil", "nil", focusStr) {
				t.Fatalf("Wrong Log.Debug focus string")
			}
			if args[6].(string) != notNilFsetStr {
				t.Fatalf("Wrong Log.Debug Fset string")
			}
			if args[7].(string) != "-" {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration", func(t *testing.T) {
		focusStr := "somePackagePath"
		config := Config{
			true,
			"abc",
			[]string{},
			token.NewFileSet(),
			[]string{},
			FocusFunctionName(focusStr),
			LogJSON,
		}
		log := &mockLoggerCLI{}
		calls := 0
		pattern := "somePattern"
		log.mockDebug = func(msgFormat string, args ...interface{}) LoggerCLI {
			calls += 1

			if msgFormat != finalConfigTemplate {
				t.Fatalf("Wrong Log.Debug template")
			}
			if len(args) != 8 {
				t.Fatalf("Wrong Log.Debug variadic args")
			}
			if args[0].(string) != pattern {
				t.Fatalf("Wrong Log.Debug pattern")
			}
			if args[1].(bool) != config.Tests {
				t.Fatalf("Wrong Log.Debug tests boolean")
			}
			if args[2].(string) != config.Dir {
				t.Fatalf("Wrong Log.Debug dir string")
			}
			if !reflect.DeepEqual(args[3].([]string), config.Env) {
				t.Fatalf("Wrong Log.Debug Env slice")
			}
			if !reflect.DeepEqual(args[4].([]string), config.BuildFlags) {
				t.Fatalf("Wrong Log.Debug BuildFlags slice")
			}
			if args[5].(string) != fmt.Sprintf(focusTemplate, "nil", "nil", "nil", "nil", focusStr) {
				t.Fatalf("Wrong Log.Debug focus string")
			}
			if args[6].(string) != notNilFsetStr {
				t.Fatalf("Wrong Log.Debug Fset string")
			}
			if args[7].(string) != "LogJSON" {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration", func(t *testing.T) {
		focusStr := "somePackagePath"
		config := Config{
			true,
			"abc",
			[]string{},
			token.NewFileSet(),
			[]string{},
			FocusFunctionName(focusStr),
			LogTrace,
		}
		log := &mockLoggerCLI{}
		calls := 0
		pattern := "somePattern"
		log.mockDebug = func(msgFormat string, args ...interface{}) LoggerCLI {
			calls += 1

			if msgFormat != finalConfigTemplate {
				t.Fatalf("Wrong Log.Debug template")
			}
			if len(args) != 8 {
				t.Fatalf("Wrong Log.Debug variadic args")
			}
			if args[0].(string) != pattern {
				t.Fatalf("Wrong Log.Debug pattern")
			}
			if args[1].(bool) != config.Tests {
				t.Fatalf("Wrong Log.Debug tests boolean")
			}
			if args[2].(string) != config.Dir {
				t.Fatalf("Wrong Log.Debug dir string")
			}
			if !reflect.DeepEqual(args[3].([]string), config.Env) {
				t.Fatalf("Wrong Log.Debug Env slice")
			}
			if !reflect.DeepEqual(args[4].([]string), config.BuildFlags) {
				t.Fatalf("Wrong Log.Debug BuildFlags slice")
			}
			if args[5].(string) != fmt.Sprintf(focusTemplate, "nil", "nil", "nil", "nil", focusStr) {
				t.Fatalf("Wrong Log.Debug focus string")
			}
			if args[6].(string) != notNilFsetStr {
				t.Fatalf("Wrong Log.Debug Fset string")
			}
			if args[7].(string) != "LogTrace" {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration", func(t *testing.T) {
		focusStr := "somePackagePath"
		config := Config{
			true,
			"abc",
			[]string{},
			token.NewFileSet(),
			[]string{},
			FocusFunctionName(focusStr),
			LogDebug,
		}
		log := &mockLoggerCLI{}
		calls := 0
		pattern := "somePattern"
		log.mockDebug = func(msgFormat string, args ...interface{}) LoggerCLI {
			calls += 1

			if msgFormat != finalConfigTemplate {
				t.Fatalf("Wrong Log.Debug template")
			}
			if len(args) != 8 {
				t.Fatalf("Wrong Log.Debug variadic args")
			}
			if args[0].(string) != pattern {
				t.Fatalf("Wrong Log.Debug pattern")
			}
			if args[1].(bool) != config.Tests {
				t.Fatalf("Wrong Log.Debug tests boolean")
			}
			if args[2].(string) != config.Dir {
				t.Fatalf("Wrong Log.Debug dir string")
			}
			if !reflect.DeepEqual(args[3].([]string), config.Env) {
				t.Fatalf("Wrong Log.Debug Env slice")
			}
			if !reflect.DeepEqual(args[4].([]string), config.BuildFlags) {
				t.Fatalf("Wrong Log.Debug BuildFlags slice")
			}
			if args[5].(string) != fmt.Sprintf(focusTemplate, "nil", "nil", "nil", "nil", focusStr) {
				t.Fatalf("Wrong Log.Debug focus string")
			}
			if args[6].(string) != notNilFsetStr {
				t.Fatalf("Wrong Log.Debug Fset string")
			}
			if args[7].(string) != "LogDebug" {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
}
