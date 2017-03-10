package ztest

import "strings"

// Bar might be one of the best functions in the world.
func Bar() string {
	message := "Hello, world"
	return strings.Repeat(message, 10)
}

func xyz() {
	Foo()
}
