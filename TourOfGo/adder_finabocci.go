package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	return func(x int) int {
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}

// func adder() func(int) int {
// 	sum := 0
// 	return func(x int) int {
// 		sum += x
// 		return sum
// 	}
// }

// func main() {
// 	pos, neg := adder(), adder()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(
// 			pos(i),
// 			neg(-2*i),
// 		)
// 	}
// }
