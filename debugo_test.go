package debugo

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

const debugEnvName = "DEBUG"

func log(debug *Debug, format string, opt ...interface{}) string {
	var buf bytes.Buffer
	debug.SetOutput(&buf)
	debug.Debugf(format, opt...)
	return buf.String()
}

func TestDebug_Debugf(t *testing.T) {
	os.Setenv(debugEnvName, "debug")
	defer os.Unsetenv(debugEnvName)
	debug := newTestDebug("debug")
	out := log(debug, "hello")

	assert.True(t, strings.Contains(out, "hello"))
	assert.True(t, strings.Contains(out, "debug"))
}

func TestDebug_Debugf2(t *testing.T) {
	os.Setenv(debugEnvName, "false")
	defer os.Unsetenv(debugEnvName)

	debug := newTestDebug("debug")
	out := log(debug, "hello")

	assert.Equal(t, "", out)
}

func TestDebug_Debugf3(t *testing.T) {
	os.Setenv(debugEnvName, "*")
	defer os.Unsetenv(debugEnvName)

	debug := newTestDebug("debug")
	out := log(debug, "hello")

	assert.True(t, strings.Contains(out, "hello"))
	assert.True(t, strings.Contains(out, "debug"))

	debug2 := newTestDebug("debug2")
	out2 := log(debug2, "hello")

	assert.True(t, strings.Contains(out2, "hello"))
	assert.True(t, strings.Contains(out2, "debug2"))

}

func TestDebug_SetColorFunc(t *testing.T) {
	os.Setenv(debugEnvName, "debug")
	defer os.Unsetenv(debugEnvName)
	debug := newTestDebug("debug")
	debug.SetColorFunc(func(s string) string {
		return s
	})
	out := log(debug, "hello")

	assert.True(t, strings.HasSuffix(out, "debug hello\n"))
}

func TestDebug_DisableTime(t *testing.T) {
	os.Setenv(debugEnvName, "debug")
	defer os.Unsetenv(debugEnvName)
	debug := newTestDebug("debug")
	debug.DisableTime()
	out := log(debug, "hello")

	assert.Equal(t, out, "\x1b[38;5;59mdebug\x1b[0;00m hello\n")
}
