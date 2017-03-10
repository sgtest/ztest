package ztest

import "fmt"

// Foo is a great function.
func Foo() int {
	return Bar()
}

func abc() {
	fmt.Println(Foo())
}
