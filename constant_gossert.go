// Copyright 2025 Francis Michael Stephens. All rights reserved.  Use of this
// source code is governed by an MIT license that can be found in the LICENSE
// file.

//go:build gossert

package gossert

// Constant determines whether the asserts are run or not.  Defaults to false.
// But if you run the compiler with the compile tag '-tags gossert' the
// runtime asserts will be activated.
const runAsserts = true
