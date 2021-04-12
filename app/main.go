package main

import (
	"fmt"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func main() {
	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)

	// person := mylib.Person{Name: "mika", Age: 45}
	// fmt.Println(person)

	// fmt.Println(mylib.Public)
	// fmt.Println(mylib.private)
	// fmt.Println(mylib.hoge)

	// s := []int{1, 2, 3, 4, 5}
	// fmt.Println(mylib.Average(s))
}

func init() {
	fmt.Println("-------------------------- init! --------------------------")
}
