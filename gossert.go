// Copyright 2025 Francis Michael Stephens. All rights reserved.  Use of this
// source code is governed by an MIT license that can be found in the LICENSE
// file.

// A small convenient package for adding runtime assertions to Go programs.
//
// By default all assertions are ignored and not executed. But if you include
// '-tags gossert' when your program is compiled the assertions will be run.
//
// When assertions are being ignored they will be eliminated by the Go compiler
// as dead code. Therefore there should close to zero performance impact for
// adding assertions when they are not used.
package gossert

import (
	"fmt"
	"os"
	"runtime/debug"
)

// Returns true if assert functions are executed in Gossert* function calls.
// False if assert functions are ignored.
//
// This function can be used to optionally perform assert setup steps which
// wouldn't be needed if the asserts are being ignored.
func WillRunAsserts() bool {
	return runAsserts
}

// If WillRunAsserts() is true, then assert is executed.
func Gossert(assert func()) {
	if runAsserts {
		assert()
	}
}

// If WillRunAsserts() is true, then assert is executed.  If assert returns a
// non-nil error then os.Exit(-1) is called.
func GossertExit(assert func() error) {
	if runAsserts {
		if err := assert(); err != nil {
			os.Exit(-1)
		}
	}
}

// If WillRunAsserts() is true, then assert is executed.  If assert returns a
// non-nil error then the error is written to stderr.
func GossertMsg(assert func() error) {
	if runAsserts {
		if err := assert(); err != nil {
			stack := debug.Stack()
			fmt.Fprintf(os.Stderr, "gossert failure: %s\n%s\n", err, stack)
		}
	}
}

// If WillRunAsserts() is true, then assert is executed.  If assert returns a
// non-nil error then the error is written to stderr and then os.Exit(-1) is
// called.
func GossertMsgExit(assert func() error) {
	if runAsserts {
		if err := assert(); err != nil {
			stack := debug.Stack()
			fmt.Fprintf(os.Stderr, "gossert failure: %s\n%s\n", err, stack)
			os.Exit(-1)
		}
	}
}
