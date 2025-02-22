package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/fmstephe/gossert"
)

var (
	flagGossertExitMsg = flag.Bool("gossert-exit-msg", false, "Demonstrates the GossertExitMsg() function")
	flagGossertExit    = flag.Bool("gossert-exit", false, "Demonstrates the GossertExit() function")
	flagGossert        = flag.Bool("gossert", false, "Demonstrates the Gossert() function")
)

// This tiny program demonstrates how we can enable
func main() {
	if gossert.WillRunAsserts() {
		fmt.Println("The asserts will be running in this example")
	} else {
		fmt.Println("The asserts are disabled and will not run. You can enable them by adding '-tags gossert' to your build command.")
	}

	flag.Parse()
	if *flagGossertExitMsg {
		demonstrateGossertExitMsg()
	}
	if *flagGossertExit {
		demonstrateGossertExit()
	}
	if *flagGossert {
		demonstrateGossert()
	}
}

func demonstrateGossertExitMsg() {
	// These will print fine
	fmt.Printf("summed to %d\n", sum(1, 2))
	fmt.Printf("summed to %d\n", sum(math.MaxInt, -math.MaxInt))

	// When asserts are enabled, this call will fail
	fmt.Printf("summed to %d\n", sum(math.MaxInt, math.MaxInt))
}

func demonstrateGossertExit() {
	// These will print fine
	fmt.Printf("subtracted to %d\n", subtract(1, 2))
	fmt.Printf("subtracted to %d\n", subtract(-math.MaxInt, math.MaxInt))

	// When asserts are enabled, this call will fail
	fmt.Printf("subtracted to %d\n", subtract(math.MaxInt, -math.MaxInt))
}

func demonstrateGossert() {
	// These will print fine
	fmt.Printf("multiplied to %d\n", multiply(1, 2))
	fmt.Printf("multiplied to %d\n", multiply(10, 20))

	// When asserts are enabled, these calls will print error messages
	fmt.Printf("multiplied to %d\n", multiply(math.MaxInt, 2))
	fmt.Printf("multiplied to %d\n", multiply(-math.MaxInt, 2))
}

// Sum two int values, assert that the result does not overflow.
//
// When asserts are enabled and overflow is detected the program will exit
func sum(x, y int) int {
	gossert.GossertExitMsg(func() (bool, string) {
		return assertSum(x, y)
	})
	return x + y
}

// Subtract two int values, assert that the result does not overflow.
//
// When asserts are enabled and overflow is detected the program will exit
func subtract(x, y int) int {
	gossert.GossertExit(func() bool {
		return assertSubtract(x, y)
	})
	return x - y
}

// Multiply two int values, assert that the result does not overflow.
//
// When asserts are enabled and overflow is detected an error message will be
// printed, but execution will continue.
func multiply(x, y int) int {
	gossert.Gossert(func() {
		assertMultiply(x, y)
	})
	return x * y
}

// A function which determines whether summing two int values will overflow It
// returns false and a detailed message on assertion failure.
func assertSum(x, y int) (bool, string) {
	sum := x + y

	// Check for overflow
	if x > 0 && sum < y {
		return false, fmt.Sprintf("%d + %d overflows to %d\n", x, y, sum)
	}

	// Check for underflow
	if x < 0 && sum > y {
		return false, fmt.Sprintf("%d + %d underflows to %d\n", x, y, sum)
	}

	return true, ""
}

// A function which determines whether subtracting two int values will
// underflow It returns false on assertion failure.
func assertSubtract(x, y int) bool {
	difference := x - y

	// Check for underflow
	if y > 0 && difference > x {
		fmt.Printf("%d - %d underflows to %d\n", x, y, difference)
		return false
	}

	// Check for overflow
	if y < 0 && difference < x {
		fmt.Printf("%d - %d overflows to %d\n", x, y, difference)
		return false
	}

	return true
}

// A function which determines whether multiplying two int values will
// overflow. It doesn't return any value, but prints an error message when it
// detects overflow.
func assertMultiply(x, y int) {
	// No overflow possible here
	if x == 0 || y == 0 {
		return
	}

	product := x * y

	if x != product/y || y != product/x {
		fmt.Printf("%d * %d overflows to %d\n", x, y, product)
	}
}
