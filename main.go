// Application which greets you.
package main

import "fmt"

//go:generate go run golang.org/x/tools/cmd/stringer -type=Country

// Country represents a country.
type Country int

const (
	Poland Country = iota // Poland represent the Republic of Poland.
)

func main() {
	fmt.Println(greet())
}

func greet() string {
	return "Hi from " + Poland.String()
}
