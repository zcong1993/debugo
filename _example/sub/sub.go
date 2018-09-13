package sub

import "github.com/zcong1993/debugo"

var debugSub = debugo.NewDebug("sub")
var debugSub2 = debugo.NewDebug("sub-sub")

// Run is runner for example
func Run() {
	debugSub.DisableTime()
	debugSub.Debugf("sub debugger %s", "output")
	for i := 0; i < 5; i++ {
		go debugSub2.Debugf("sub2 debugger %s %d", "output2", i)
	}
}
