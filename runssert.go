package gossert

import (
	"fmt"
	"os"
	"runtime/debug"
)

func Gossert(assert func()) {
	if runAsserts {
		assert()
	}
}

func GossertExit(assert func() bool) {
	if runAsserts {
		ok := assert()
		if !ok {
			stack := debug.Stack()
			fmt.Fprintf(os.Stderr, "runssert failure: %s\n", stack)
			os.Exit(-1)
		}
	}
}

func GossertExitMsg(assert func() (bool, string)) {
	if runAsserts {
		ok, msg := assert()
		if !ok {
			stack := debug.Stack()
			fmt.Fprintf(os.Stderr, "runssert failure: %s\n%s\n", msg, stack)
			os.Exit(-1)
		}
	}
}
