package main

import (
	. "Framework/Frame"
	"fmt"
	// "strconv"
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
	// Framework.InitDebug()
	// Framework.RunDebug()
	Framework.Init()
	Framework.Run()
	Framework.Close()
	fmt.Println("Main package finished running.")
}
