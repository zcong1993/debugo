# debugo [![Go Report Card](https://goreportcard.com/badge/github.com/zcong1993/debugo)](https://goreportcard.com/report/github.com/zcong1993/debugo)

> debug for golang

## Install

```bash
$ go get -v github.com/zcong1993/debugo
```

## Usage

see [example](./example)

```go
package main

import (
	"github.com/zcong1993/debugo"
)

var debugMain = debugo.NewDebug("main")

func main() {
	debugMain.Debugf("debug main")
}
```

## License

MIT &copy; zcong1993
