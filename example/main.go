package main

import (
	"github.com/zcong1993/debugo"
	"github.com/zcong1993/debugo/example/sub"
)

var debugMain = debugo.NewDebug("main")

func main() {
	debugMain.Debugf("debug main")
	sub.Run()
}
