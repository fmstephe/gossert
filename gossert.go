package gossert

import (
	"fmt"
	"os"
	"runtime/debug"
)

func WillRunAsserts() bool {
	return runAsserts
}

func Gossert(assert func()) {
	if runAsserts {
		assert()
	}
}

func GossertExit(assert func() error) {
	if runAsserts {
		if err := assert(); err != nil {
			os.Exit(-1)
		}
	}
}

func GossertMsg(assert func() error) {
	if runAsserts {
		if err := assert(); err != nil {
			stack := debug.Stack()
			fmt.Fprintf(os.Stderr, "gossert failure: %s\n%s\n", err, stack)
		}
	}
}

func GossertExitMsg(assert func() error) {
	if runAsserts {
		if err := assert(); err != nil {
			stack := debug.Stack()
			fmt.Fprintf(os.Stderr, "gossert failure: %s\n%s\n", err, stack)
			os.Exit(-1)
		}
	}
}
