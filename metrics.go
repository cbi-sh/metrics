package metrics

import (
	"fmt"
	"sync/atomic"
)

var code200 uint64
var code400 uint64
var code404 uint64
var code500 uint64

func Inc200() {
	atomic.AddUint64(&code200, 1)
}

func Inc400() {
	atomic.AddUint64(&code400, 1)
}

func Inc404() {
	atomic.AddUint64(&code404, 1)
}

func Inc500() {
	atomic.AddUint64(&code500, 1)
}

func Get() string {
	return fmt.Sprint(
		"code200 ", code200, "\n",
		"code400 ", code400, "\n",
		"code404 ", code404, "\n",
		"code500 ", code500, "\n",
	)
}
