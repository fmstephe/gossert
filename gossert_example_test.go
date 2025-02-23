package gossert

import (
	"fmt"
	"os"
)

func Example_Gossert() {
	// Imagine these are function args being passed into a function
	x := 10
	y := -10

	// By default the asserts are switched off
	// But if we run these tests with '-tags gossert' the assertion will trigger a panic
	//
	// When using Gossert we take full control with how failed assertions are handled and reported
	Gossert(func() {
		if x < 0 {
			err := fmt.Errorf("x is negative %d\n", x)
			fmt.Fprintf(os.Stderr, "%s", err)
			panic(err)
		}
		if y < 0 {
			err := fmt.Errorf("y is negative %d\n", y)
			fmt.Fprintf(os.Stderr, "%s", err)
			panic(err)
		}
	})

	// If we reach here and are able to perform some computation on the args
	fmt.Printf("%d\n", x+y)
	// Output: 0
}

func Example_GossertMsg() {
	// Imagine these are function args being passed into a function
	x := 10
	y := -10

	// By default the asserts are switched off
	// But if we run these tests with '-tags gossert' the assertion will trigger a panic
	//
	// When using GossertMsg we return an error on an assertion failure
	//
	// A non-nil error will be printed to stderr along with a stack trace
	GossertMsg(func() error {
		if x < 0 {
			return fmt.Errorf("x is negative %d\n", x)
		}
		if y < 0 {
			return fmt.Errorf("y is negative %d\n", y)
		}

		return nil
	})

	// If we reach here and are able to perform some computation on the args
	fmt.Printf("%d\n", x+y)
	// Output: 0
}

func Example_GossertExit() {
	// Imagine these are function args being passed into a function
	x := 10
	y := -10

	// By default the asserts are switched off
	// But if we run these tests with '-tags gossert' the assertion will trigger a panic
	//
	// When using GossertExit we return an error on an assertion failure
	//
	// We must take responsibility for logging the failure, a non-nil error
	// will cause os.Exit(-1) to be called inside GossertExit()
	GossertExit(func() error {
		if x < 0 {
			err := fmt.Errorf("x is negative %d\n", x)
			fmt.Fprintf(os.Stderr, "%s", err)
			return err
		}
		if y < 0 {
			err := fmt.Errorf("y is negative %d\n", y)
			fmt.Fprintf(os.Stderr, "%s", err)
			return err
		}

		return nil
	})

	// If we reach here and are able to perform some computation on the args
	fmt.Printf("%d\n", x+y)
	// Output: 0
}

func Example_GossertMsgExit() {
	// Imagine these are function args being passed into a function
	x := 10
	y := -10

	// By default the asserts are switched off
	// But if we run these tests with '-tags gossert' the assertion will trigger a panic
	//
	// When using GossertMsgExit we return an error on an assertion failure
	//
	// A non-nil error will be printed to stderr along with a stack trace
	// and os.Exit(-1) will be called inside GossertMsgExit()
	GossertMsgExit(func() error {
		if x < 0 {
			return fmt.Errorf("x is negative %d\n", x)
		}
		if y < 0 {
			return fmt.Errorf("y is negative %d\n", y)
		}

		return nil
	})

	// If we reach here and are able to perform some computation on the args
	fmt.Printf("%d\n", x+y)
	// Output: 0
}
