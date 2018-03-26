package ztest

import "fmt"
import "github.com/gorilla/mux.git"

var _ = mux.NewRouter

// Foo is a great function.
func Foo() int {
	fmt.Println(Bar())
	abc := 123
	return abc
}

func abc() {
	fmt.Println(Foo())
}
