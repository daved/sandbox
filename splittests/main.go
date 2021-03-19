// +build tray

// Splittests demonstrates that we can hide the main file/func for testing
// and still be able to build the intended application.
package main

import (
	"github.com/thisdoesnotexist/stardust"
)

func main() {
	stardust.CooCaChoo("google me")
}
