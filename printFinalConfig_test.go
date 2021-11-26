package parser

import (
	"fmt"
	"github.com/mathbalduino/go-log/loggerCLI"
	"go/token"
	"reflect"
	"testing"
)

func TestPrintFinalConfig(t *testing.T) {
	t.Run("Print the correct final configuration when focus, fset, dir and logger are nil", func(t *testing.T) {
		config := Config{
			true,
			"",
			[]string{},
			nil,
			[]string{},
			nil,
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
			if len(args) != 9 {
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
			if args[7].(string) != nilLoggerStr {
				t.Fatalf("Wrong Log.Debug Logger string")
			}
			if args[8].(string) != emptyLogFlagsStr {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration when fset, dir and logger are nil (focus pkg)", func(t *testing.T) {
		focusStr := "somePackagePath"
		config := Config{
			true,
			"",
			[]string{},
			nil,
			[]string{},
			FocusPackagePath(focusStr),
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
			if len(args) != 9 {
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
			if args[5].(string) != fmt.Sprintf(focusTemplate, focusStr, "nil", "nil") {
				t.Fatalf("Wrong Log.Debug focus string")
			}
			if args[6].(string) != nilFsetStr {
				t.Fatalf("Wrong Log.Debug Fset string")
			}
			if args[7].(string) != nilLoggerStr {
				t.Fatalf("Wrong Log.Debug Logger string")
			}
			if args[8].(string) != emptyLogFlagsStr {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration when fset, dir and logger are nil (focus filepath)", func(t *testing.T) {
		filepath := "someFilepath"
		config := Config{
			true,
			"",
			[]string{},
			nil,
			[]string{},
			FocusFilePath(filepath),
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
			if len(args) != 9 {
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
			if args[5].(string) != fmt.Sprintf(focusTemplate, "nil", filepath, "nil") {
				t.Fatalf("Wrong Log.Debug focus string")
			}
			if args[6].(string) != nilFsetStr {
				t.Fatalf("Wrong Log.Debug Fset string")
			}
			if args[7].(string) != nilLoggerStr {
				t.Fatalf("Wrong Log.Debug Logger string")
			}
			if args[8].(string) != emptyLogFlagsStr {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration when fset, dir and logger are nil (focus typeName)", func(t *testing.T) {
		typeName := "someTypeName"
		config := Config{
			true,
			"",
			[]string{},
			nil,
			[]string{},
			FocusTypeName(typeName),
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
			if len(args) != 9 {
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
			if args[5].(string) != fmt.Sprintf(focusTemplate, "nil", "nil", typeName) {
				t.Fatalf("Wrong Log.Debug focus string")
			}
			if args[6].(string) != nilFsetStr {
				t.Fatalf("Wrong Log.Debug Fset string")
			}
			if args[7].(string) != nilLoggerStr {
				t.Fatalf("Wrong Log.Debug Logger string")
			}
			if args[8].(string) != emptyLogFlagsStr {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration when dir and logger are nil", func(t *testing.T) {
		typeName := "someTypeName"
		config := Config{
			true,
			"",
			[]string{},
			token.NewFileSet(),
			[]string{},
			FocusTypeName(typeName),
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
			if len(args) != 9 {
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
			if args[5].(string) != fmt.Sprintf(focusTemplate, "nil", "nil", typeName) {
				t.Fatalf("Wrong Log.Debug focus string")
			}
			if args[6].(string) != notNilFsetStr {
				t.Fatalf("Wrong Log.Debug Fset string")
			}
			if args[7].(string) != nilLoggerStr {
				t.Fatalf("Wrong Log.Debug Logger string")
			}
			if args[8].(string) != emptyLogFlagsStr {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration when logger is not nil", func(t *testing.T) {
		typeName := "someTypeName"
		dir := "someDir"
		config := Config{
			true,
			dir,
			[]string{},
			token.NewFileSet(),
			[]string{},
			FocusTypeName(typeName),
			loggerCLI.New(false, 0),
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
			if len(args) != 9 {
				t.Fatalf("Wrong Log.Debug variadic args")
			}
			if args[0].(string) != pattern {
				t.Fatalf("Wrong Log.Debug pattern")
			}
			if args[1].(bool) != config.Tests {
				t.Fatalf("Wrong Log.Debug tests boolean")
			}
			if args[2].(string) != dir {
				t.Fatalf("Wrong Log.Debug dir string")
			}
			if !reflect.DeepEqual(args[3].([]string), config.Env) {
				t.Fatalf("Wrong Log.Debug Env slice")
			}
			if !reflect.DeepEqual(args[4].([]string), config.BuildFlags) {
				t.Fatalf("Wrong Log.Debug BuildFlags slice")
			}
			if args[5].(string) != fmt.Sprintf(focusTemplate, "nil", "nil", typeName) {
				t.Fatalf("Wrong Log.Debug focus string")
			}
			if args[6].(string) != notNilFsetStr {
				t.Fatalf("Wrong Log.Debug Fset string")
			}
			if args[7].(string) != notNilLoggerStr {
				t.Fatalf("Wrong Log.Debug Logger string")
			}
			if args[8].(string) != ignoredLogFlagsStr {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration, with LogJSON", func(t *testing.T) {
		config := Config{
			true,
			"abc",
			[]string{},
			token.NewFileSet(),
			[]string{},
			nil,
			nil,
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
			if len(args) != 9 {
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
			if args[7].(string) != nilLoggerStr {
				t.Fatalf("Wrong Log.Debug Logger string")
			}
			if args[8].(string) != "LogJSON" {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration, with LogTrace", func(t *testing.T) {
		config := Config{
			true,
			"abc",
			[]string{},
			token.NewFileSet(),
			[]string{},
			nil,
			nil,
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
			if len(args) != 9 {
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
			if args[7].(string) != nilLoggerStr {
				t.Fatalf("Wrong Log.Debug Logger string")
			}
			if args[8].(string) != "LogTrace" {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration, with LogDebug", func(t *testing.T) {
		config := Config{
			true,
			"abc",
			[]string{},
			token.NewFileSet(),
			[]string{},
			nil,
			nil,
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
			if len(args) != 9 {
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
			if args[7].(string) != nilLoggerStr {
				t.Fatalf("Wrong Log.Debug Logger string")
			}
			if args[8].(string) != "LogDebug" {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration, with LogInfo", func(t *testing.T) {
		config := Config{
			true,
			"abc",
			[]string{},
			token.NewFileSet(),
			[]string{},
			nil,
			nil,
			LogInfo,
		}
		log := &mockLoggerCLI{}
		calls := 0
		pattern := "somePattern"
		log.mockDebug = func(msgFormat string, args ...interface{}) LoggerCLI {
			calls += 1

			if msgFormat != finalConfigTemplate {
				t.Fatalf("Wrong Log.Debug template")
			}
			if len(args) != 9 {
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
			if args[7].(string) != nilLoggerStr {
				t.Fatalf("Wrong Log.Debug Logger string")
			}
			if args[8].(string) != "LogInfo" {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration, with LogWarn", func(t *testing.T) {
		config := Config{
			true,
			"abc",
			[]string{},
			token.NewFileSet(),
			[]string{},
			nil,
			nil,
			LogWarn,
		}
		log := &mockLoggerCLI{}
		calls := 0
		pattern := "somePattern"
		log.mockDebug = func(msgFormat string, args ...interface{}) LoggerCLI {
			calls += 1

			if msgFormat != finalConfigTemplate {
				t.Fatalf("Wrong Log.Debug template")
			}
			if len(args) != 9 {
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
			if args[7].(string) != nilLoggerStr {
				t.Fatalf("Wrong Log.Debug Logger string")
			}
			if args[8].(string) != "LogWarn" {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration, with LogError", func(t *testing.T) {
		config := Config{
			true,
			"abc",
			[]string{},
			token.NewFileSet(),
			[]string{},
			nil,
			nil,
			LogError,
		}
		log := &mockLoggerCLI{}
		calls := 0
		pattern := "somePattern"
		log.mockDebug = func(msgFormat string, args ...interface{}) LoggerCLI {
			calls += 1

			if msgFormat != finalConfigTemplate {
				t.Fatalf("Wrong Log.Debug template")
			}
			if len(args) != 9 {
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
			if args[7].(string) != nilLoggerStr {
				t.Fatalf("Wrong Log.Debug Logger string")
			}
			if args[8].(string) != "LogError" {
				t.Fatalf("Wrong Log.Debug LogFlags string")
			}
			return log
		}
		printFinalConfig(pattern, config, log)
		if calls != 1 {
			t.Fatalf("Log.Debug was expected to be called one time")
		}
	})
	t.Run("Print the correct final configuration, with LogFatal", func(t *testing.T) {
		config := Config{
			true,
			"abc",
			[]string{},
			token.NewFileSet(),
			[]string{},
			nil,
			nil,
			LogFatal,
		}
		log := &mockLoggerCLI{}
		calls := 0
		pattern := "somePattern"
		log.mockDebug = func(msgFormat string, args ...interface{}) LoggerCLI {
			calls += 1

			if msgFormat != finalConfigTemplate {
				t.Fatalf("Wrong Log.Debug template")
			}
			if len(args) != 9 {
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
			if args[7].(string) != nilLoggerStr {
				t.Fatalf("Wrong Log.Debug Logger string")
			}
			if args[8].(string) != "LogFatal" {
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
		config := Config{
			true,
			"abc",
			[]string{},
			token.NewFileSet(),
			[]string{},
			nil,
			nil,
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
			if len(args) != 9 {
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
			if args[7].(string) != nilLoggerStr {
				t.Fatalf("Wrong Log.Debug Logger string")
			}
			if args[8].(string) != "LogDebug" {
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
