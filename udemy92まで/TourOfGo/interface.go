package main

import "fmt"

type I interface {
	M()
}

type In interface {
	Hoge()
	M()
}

type T struct {
	S string
}

type St struct {
	I int
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t St) M() {
	fmt.Println(t.S)
}

func (st St) Hoge() {
	fmt.Println(st.I)
}

func main() {
	// var i I = T{"hello"}
	// i.M()

	var in In = St{1, "hoge"}
	in.Hoge()
	in.M()
}





package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I
	var t = &T{"hoge"}
	i = t
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}




package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type PerfectHuman struct {
	FirstName string
	LastName  string
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func (pf PerfectHuman) Hoge() string {
	return fmt.Sprintf("%v %v です！", pf.LastName, pf.FirstName)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	pf := PerfectHuman{"Atsuhiko", "Nakata"}
	fmt.Println(pf.FirstName, pf.LastName)
	fmt.Println(pf.Hoge())
}
