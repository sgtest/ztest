package ztest

// Qux is an awesome type.
type Qux struct {
	AAA string
	BBB int
	CCC *Qux
}

// Zip is an even better function.
func Zip() *Qux {
	Foo()
	Bar()
	return &Qux{
		AAA: "hello",
		BBB: 777,
		CCC: &Qux{AAA: "world"},
	}
}
