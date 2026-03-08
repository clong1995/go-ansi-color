package pcolor

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestColor(t *testing.T) {
	tests := []struct {
		name  string
		color int
		str   string
		want  string
	}{
		{"Red", 31, "hello", "\x1b[31mhello\x1b[0m"},
		{"Green", 32, "world", "\x1b[32mworld\x1b[0m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Color(tt.color, tt.str); got != tt.want {
				t.Errorf("Color() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestFormatFunctions(t *testing.T) {
	tests := []struct {
		name     string
		function func(string, ...any) string
		format   string
		args     []any
		want     string
	}{
		{"Err", Err, "error %d", []any{1}, "\x1b[31m❌ error 1\x1b[0m"},
		{"Succ", Succ, "success %s", []any{"abc"}, "\x1b[32m✅ success abc\x1b[0m"},
		{"Warn", Warn, "warning", nil, "\x1b[33m⚠️ warning\x1b[0m"},
		{"Fatal", Fatal, "fatal error", nil, "\x1b[91m💀 fatal error\x1b[0m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.function(tt.format, tt.args...); got != tt.want {
				t.Errorf("%s() = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}

func TestError(t *testing.T) {
	err := fmt.Errorf("test error")
	want := "\x1b[31m❌ test error\x1b[0m"
	if got := Error(err); got != want {
		t.Errorf("Error() = %q, want %q", got, want)
	}
}

func TestPrintFunctions(t *testing.T) {
	var outBuf, errBuf bytes.Buffer
	outLogger = log.New(&outBuf, "", 0)
	errLogger = log.New(&errBuf, "", 0)

	// Restore original loggers after test
	defer func() {
		outLogger = log.New(os.Stdout, "", 0)
		errLogger = log.New(os.Stderr, "", 0)
	}()

	tests := []struct {
		name     string
		function func(string, string, ...any)
		prefix   string
		format   string
		args     []any
		wantBuf  *bytes.Buffer
		want     string
	}{
		{"PrintSucc", PrintSucc, "test", "success", nil, &outBuf, "\x1b[1m[test]\x1b[0m \x1b[32m✅ success\x1b[0m\n"},
		{"PrintError", PrintError, "test", "error", nil, &errBuf, "\x1b[1m[test]\x1b[0m \x1b[31m❌ error\x1b[0m\n"},
		{"PrintWarn", PrintWarn, "test", "warning", nil, &errBuf, "\x1b[1m[test]\x1b[0m \x1b[33m⚠️ warning\x1b[0m\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantBuf.Reset()
			tt.function(tt.prefix, tt.format, tt.args...)
			if got := tt.wantBuf.String(); got != tt.want {
				t.Errorf("%s() output = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}

func TestPrintFatal(t *testing.T) {
	if os.Getenv("GO_TEST_FATAL") == "1" {
		PrintFatal("test", "fatal error")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestPrintFatal")
	cmd.Env = append(os.Environ(), "GO_TEST_FATAL=1")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()

	if e, ok := errors.AsType[*exec.ExitError](err); !ok || e.Success() {
		t.Fatalf("process ran with err %v, want exit status 1", err)
	}

	want := "\x1b[1m[test]\x1b[0m \x1b[91m💀 fatal error\x1b[0m\n"
	// The log package adds a timestamp, so we check if the output contains our desired string
	if got := stderr.String(); !strings.HasSuffix(got, want) {
		t.Errorf("PrintFatal() output = %q, want suffix %q", got, want)
	}
}
