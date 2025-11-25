// Application which greets you.
package main

import "fmt"

//go:generate go run golang.org/x/tools/cmd/stringer -type=Country

// Country represents a country.
type Country int

const (
	// Poland represent the Republic of Poland.
	Poland Country = iota
)

func main() {
	fmt.Println(greet())
}

func greet() string {
	return "Hi from " + Poland.String()
}
