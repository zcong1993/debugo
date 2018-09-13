# debugo [![Go Report Card](https://goreportcard.com/badge/github.com/zcong1993/debugo)](https://goreportcard.com/report/github.com/zcong1993/debugo) [![CircleCI branch](https://img.shields.io/circleci/project/github/zcong1993/debugo/master.svg)](https://circleci.com/gh/zcong1993/debugo/tree/master) [![codecov](https://codecov.io/gh/zcong1993/debugo/branch/master/graph/badge.svg)](https://codecov.io/gh/zcong1993/debugo)

> debug for golang

## Install

```bash
$ go get -v github.com/zcong1993/debugo
```

## Usage

see [example](./_example)

```go
package main

import (
	"github.com/zcong1993/debugo"
)

var debugMain = debugo.NewDebug("main")

// DEBUG=main go run main.go
func main() {
	debugMain.Debugf("debug main")
}
```

## License

MIT &copy; zcong1993
