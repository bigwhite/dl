// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The go1.11beta3 command runs the go command from go1.11beta3.
//
// To install, run:
//
//     $ go get golang.org/dl/go1.11beta3
//     $ go1.11beta3 download
//
// And then use the go1.11beta3 command as if it were your normal go
// command.
//
// See the release notes at https://golang.org/doc/go1.11beta3
//
// File bugs at https://golang.org/issues/new
package main

import "golang.org/dl/internal/version"

func main() {
	version.Run("go1.11beta3")
}
