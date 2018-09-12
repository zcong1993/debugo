package debugo

import (
	"fmt"
	color "github.com/aybabtme/rgbterm"
	"io"
	"os"
	"strings"
	"time"
)

// ColorFunc is color function
type ColorFunc func(string) string

// Debug is struct of debug
type Debug struct {
	Out       io.Writer
	name      string
	colorFunc ColorFunc
	show      bool
	showTime  bool
}

const (
	// ALL is wildcard string
	ALL = "*"
	// FALSE is disable string
	FALSE = "false"
	// TimeFormat is our time format layout
	TimeFormat = "2006-01-02 15:04:05"
)

var cachedEnv []string
var isInit = false

func emptyColorFunc(in string) string {
	return ""
}

func genRgbColorFunc(r, g, b uint8) ColorFunc {
	return func(s string) string {
		return color.FgString(s, r, g, b)
	}
}

func parseEnv() []string {
	var res []string
	str := os.Getenv("DEBUG")
	tmpArr := strings.Split(str, ",")
	for _, v := range tmpArr {
		vv := strings.TrimSpace(v)
		if vv == ALL {
			return []string{"*"}
		}
		if vv == FALSE {
			return []string{}
		}
		res = append(res, vv)
	}
	return res
}

func matchName(name string, env []string) bool {
	for _, v := range env {
		if v == ALL {
			return true
		}
		if v == FALSE {
			return false
		}
		if name == v {
			return true
		}
	}
	return false
}

// NewDebug return a new debug instance
func NewDebug(name string) *Debug {
	if !isInit {
		cachedEnv = parseEnv()
		isInit = true
	}

	n := strings.TrimSpace(name)
	show := matchName(n, cachedEnv)
	colorFunc := emptyColorFunc
	if show {
		r, g, b := genColor([]byte(n))
		colorFunc = genRgbColorFunc(r, g, b)
	}

	return &Debug{
		Out:       os.Stdout,
		name:      n,
		show:      show,
		colorFunc: colorFunc,
		showTime:  true,
	}
}

// SetOutput set an io.Writer as output, default is os.Stdout
func (debug *Debug) SetOutput(out io.Writer) {
	debug.Out = out
}

// SetColorFunc set a custom color function
func (debug *Debug) SetColorFunc(colorFunc ColorFunc) {
	debug.colorFunc = colorFunc
}

// DisableTime let logger not show current time
func (debug *Debug) DisableTime() {
	debug.showTime = false
}

// Debugf try to log debug message
func (debug *Debug) Debugf(format string, opts ...interface{}) {
	if debug.show {
		var formatStr string
		if debug.showTime {
			t := time.Now().Format(TimeFormat)
			formatStr = debug.colorFunc(fmt.Sprintf("%s %s ", t, debug.name)) + format + "\n"
		} else {
			formatStr = debug.colorFunc(debug.name) + " " + format + "\n"
		}
		fmt.Fprintf(debug.Out, formatStr, opts...)
	}
}
