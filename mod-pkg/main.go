package main

import (
	"fmt"
	"mod-pkg/bar"
	"mod-pkg/foo"
)

// Show return main
func Show() string {
	return "main"
}

func main() {
	fmt.Println("hello from", Show())
	fmt.Println("hello from", bar.Show())
	fmt.Println("hello from", foo.Show())
}
