package parser

import (
	"github.com/mathbalduino/go-log/loggerCLI"
	"go/token"
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	t.Run("If not provided, FileSet should be created", func(t *testing.T) {
		config := Config{Fset: nil}
		packagesConfig := packagesLoadConfig(config, loggerCLI.New(true, LogDebug | LogTrace))
		if packagesConfig.Fset == nil {
			t.Fatal("Fset was expected to be not nil")
		}
	})
	t.Run("The PackagesMode should be equal to the defined constant", func(t *testing.T) {
		packagesConfig := packagesLoadConfig(Config{}, loggerCLI.New(true, LogDebug | LogTrace))
		if packagesConfig.Mode != packagesConfigMode {
			t.Fatal("The mode is not set correctly")
		}
	})
	t.Run("Packages Logf should call LogCLI Debug method, forwarding all args", func(t *testing.T) {
		logMsg := "abc"
		logArgs := []interface{}{"1", 2, 5.4}
		debugCalls := 0
		mock := &mockLoggerCLI{}
		mock.mockDebug = func(msgFormat string, args ...interface{}) LoggerCLI {
			debugCalls += 1
			if msgFormat != logMsg {
				t.Fatalf("LogCLI Debug msgformat should be equal to the one given to Logf")
			}
			if !reflect.DeepEqual(args, logArgs) {
				t.Fatalf("LogCLI Debug variadic args should be equal to the one given to Logf")
			}
			return mock
		}

		packagesConfig := packagesLoadConfig(Config{}, mock)
		packagesConfig.Logf(logMsg, logArgs...)
		if debugCalls != 1 {
			t.Fatalf("LogCLI Debug was expected to be called one time")
		}
	})
	t.Run("PackagesConfig Context, ParseFile and Overlay should be always nil", func(t *testing.T) {
		packagesConfig := packagesLoadConfig(Config{}, loggerCLI.New(true, LogDebug | LogTrace))
		if packagesConfig.Context != nil {
			t.Fatalf("Context was expected to be nil")
		}
		if packagesConfig.ParseFile != nil {
			t.Fatalf("ParseFile was expected to be nil")
		}
		if packagesConfig.Overlay != nil {
			t.Fatalf("Overlay was expected to be nil")
		}
	})
	t.Run("PackagesConfig Env, BuildFlags, Tests, Dir and Fset should be customizable", func(t *testing.T) {
		config := Config{
			Tests:      true,
			Dir:        "testeDir",
			Env:        []string{"a", "b", "c"},
			BuildFlags: []string{"1", "2", "3"},
			Fset:       token.NewFileSet(),
		}
		packagesConfig := packagesLoadConfig(config, loggerCLI.New(true, LogDebug | LogTrace))
		if packagesConfig.Tests != config.Tests {
			t.Fatalf("PackagesConfig.Tests was expected to be equal to config.Tests")
		}
		if packagesConfig.Dir != config.Dir {
			t.Fatalf("PackagesConfig.Dir was expected to be equal to config.Dir")
		}
		if !reflect.DeepEqual(packagesConfig.Env, config.Env) {
			t.Fatalf("PackagesConfig.Env was expected to be equal to config.Env")
		}
		if !reflect.DeepEqual(packagesConfig.BuildFlags, config.BuildFlags) {
			t.Fatalf("PackagesConfig.BuildFlags was expected to be equal to config.BuildFlags")
		}
		if packagesConfig.Fset != config.Fset {
			t.Fatalf("PackagesConfig.Fset was expected to be equal to config.Fset")
		}
	})
}
