package Main

import (
	. "Framework"
	"fmt"
)

// This gets whenever the containing package is imported.
// Import order:
// init() of imported packages
// package level variables initialized
// this init()
// main()

func init() {
	fmt.Println("Main package initialized.")
}

func main() {
	fmt.Println("Main package running...")
	Framework := Framework{}
	Framework.Run()
}
