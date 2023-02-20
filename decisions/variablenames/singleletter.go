package variablenames

import (
	"fmt"
	"io"
	"net/http"
)

//Single-letter variable names can be a useful tool to minimize repetition, but can also make code needlessly opaque.
//Limit their use to instances where the full word is obvious and where it would be repetitive for it to appear in
//place of the single-letter variable.

// * For a method receiver variable, a one-letter or two-letter name is preferred.
func (sw SuperWriter) Write() error {
	return nil
}

type SuperWriter struct{}

// * Using familiar variable names for common types is often helpful:
func ioRead(r io.Reader, w io.Writer)                {}
func httpRead(r http.Request, w http.ResponseWriter) {}

// * Single-letter identifiers are acceptable as integer loop variables, particularly for indices (e.g., i) and coordinates (e.g., x and y).
func someFunc() {
	for i, v := range []int{} {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}
}

// * Abbreviations can be acceptable loop identifiers when the scope is short, for example for _, n := range nodes { ... }.
func short() {
	var users []string

	for _, u := range users {
		fmt.Printf("user: %s\n", u)
	}
}
