package main

import "fmt"

type I interface { // https://go.dev/ref/spec#Type_definitions
	M()
}

type T struct { // composite type definition
	S string
}

func (t *T) M() { // method
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T // just declared var t, it's empty
	i = t // t is empty but we assign it to i 
	describe(i)
	i.M() // i is still nil, but we call the M method on it, that it got from t

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
