package ztest

import "fmt"

// Foo is a great function.
func Foo() int {
	fmt.Println(Bar())
	abc := 123
	return abc
}

func abc() {
	fmt.Println(Foo())
}
