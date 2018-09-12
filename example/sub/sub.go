package sub

import "github.com/zcong1993/debugo"

var debugSub = debugo.NewDebug("sub")

// Run is runner for example
func Run() {
	debugSub.DisableTime()
	debugSub.Debugf("sub debugger %s", "output")
}
