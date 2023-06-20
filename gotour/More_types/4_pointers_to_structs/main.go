package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	py := &v.Y
	p.X = 123e3
	fmt.Println(*py, v)
}
