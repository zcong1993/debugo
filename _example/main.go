package main

import (
	"github.com/zcong1993/debugo"
	"github.com/zcong1993/debugo/example/sub"
	"time"
)

var debugMain = debugo.NewDebug("main")

// DEBUG=main go run main.go
// DEBUG=sub go run main.go
// DEBUG=sub* go run main.go
// DEBUG=* go run main.go
// DEBUG=false go run main.go
func main() {
	debugMain.Debugf("debug main")
	sub.Run()
	time.Sleep(time.Millisecond * 100)
}
