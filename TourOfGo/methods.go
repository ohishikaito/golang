package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// func (v Vertex) Abs() float64 {
// 	return v.X + v.Y
// 	// return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(v.Abs())
// }

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) Hoge() float64 {
	return v.X + v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Hoge())
	fmt.Println(Abs(v))
}

package main

import (
	"fmt"
	"math"
)

type MyFloat float64
type Hoge string

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (s Hoge) Fuga() string {
	fmt.Println(s, "中だよ")
	return string(s + "DS")
}

func main() {
	f := MyFloat(-math.Sqrt2)
	s := Hoge("hoge")
	fmt.Println(s, "s")
	fmt.Println(s.Fuga())
	fmt.Println(f.Abs())
}




package main

import "fmt"

func main() {
	var inter_face interface{} = 4545
	int1, okok := inter_face.(int)
	fmt.Println(int1, okok)

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
