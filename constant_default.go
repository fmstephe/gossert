//go:build !runssert

package gossert

// Constant determines whether the asserts are run or not.  Defaults to false.
// But if you run the compiler with the compile tag '-tags runssert' the
// runtime asserts will be activated.
const runAsserts = false
