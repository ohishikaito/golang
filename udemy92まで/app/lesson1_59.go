package main

import (
	"fmt"
)

// func add(x int, y int) (int, int) {
// 	return x + y, x - y
// }

// func add2(s string) string {
// 	return s
// }

// func cal(price, item int) (result int) {
// 	result = price * item
// 	return
// }

// func ig() func() int {
// 	x := 10
// 	return func() (n int) {
// 		x++
// 		n += x
// 		return
// 	}
// }

// func ca(pi float64) func(radius float64) float64 {
// 	return func(radius float64) float64 {
// 		return pi * radius * radius
// 	}
// }

// func foo(params ...int) {
// 	// fmt.Println(params)
// 	for a, param := range params {
// 		fmt.Println(a, param)
// 	}
// }

// func panik() {
// 	defer func() {
// 		s := recover()
// 		fmt.Println(s)
// 	}()
// 	panic("hoge")
// }

// func errors() {
// 	if _, err := os.Open("hoge"); err != nil {
// 		fmt.Println(err)
// 	}
// }

// func calc(m map[string]int) (result int) {
// 	for _, num := range m {
// 		result += num
// 	}
// 	return
// }

// func one(x *int) (res int) {
// 	*x = 1
// 	res = 3
// 	return
// }

// type Person struct {
// 	Name string
// 	Age  int
// }

// type Vertex struct {
// 	X, Y int
// 	S    string
// }

// func c1(v Vertex) {
// 	v.X = 10
// }

// func c2(v *Vertex) {
// 	v.X = 10
// }

// type Vertex struct {
// 	x, y int
// }

// func New(x, y int) *Vertex {
// 	return &Vertex{x, y}
// }

// func (v Vertex) Area() int {
// 	return v.x * v.y
// }

// func (v *Vertex) Scale(i int) {
// 	v.x = v.x * i
// 	v.y = v.y * i
// }

// ----------------------------------------
// type Vertex3D struct {
// 	Vertex
// 	z int
// }

// func New3D(x, y, z int) *Vertex3D {
// 	return &Vertex3D{Vertex{x, y}, z}
// }

// func (v Vertex3D) Area3D() int {
// 	return v.x * v.y * v.z
// }

// func (v *Vertex3D) Scale3D(i int) {
// 	v.x = v.x * i
// 	v.y = v.y * i
// 	v.z = v.z * i
// }

// type MyInt int

// func (i MyInt) Double() int {
// 	return int(i * 2)
// }

// type Human interface {
// 	Say() string
// }

// type Person struct {
// 	Name string
// }

// type Dog struct {
// 	Name string
// }

// func (p *Person) Say() string {
// 	p.Name = "Mr." + p.Name
// 	fmt.Println(p.Name)
// 	return p.Name
// }

// func DriveCar(human Human) {
// 	if human.Say() == "Mr.Mike" {
// 		fmt.Println(true)
// 	} else {
// 		fmt.Println(false)
// 	}
// }

// func Naku(dog Dog) string {
// 	// fmt.Println(dog.Name)
// 	return dog.Name
// }

// func do(i interface{}) {
// 	// ii := i.(int)
// 	// ii *= 2
// 	// fmt.Println(ii)

// 	// ss := i.(string)
// 	// fmt.Println(ss + "!")

// 	switch v := i.(type) {
// 	case int:
// 		fmt.Println(v * 1000)
// 	case string:
// 		fmt.Println(v + "!!!!")
// 	}
// }

// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) String() string {
// 	return fmt.Sprintf("%v!!!!!!", p.Name)
// }

// type UserNotFound struct {
// 	Username string
// }

// func (e *UserNotFound) Error() (err string) {
// 	err = e.Username + "がエラーです！"
// 	return
// }

// func myFunc() error {
// 	ok := false
// 	if ok {
// 		return nil
// 	}
// 	return &UserNotFound{Username: "mike"}
// }

// type Vertex struct {
// 	X, Y int
// }

// func (v Vertex) Plus() (result int) {
// 	result = v.X + v.Y
// 	return
// }

// func (v Vertex) String() string {
// 	// return fmt.Sprintf("%v ", v.X, v.Y)
// 	return fmt.Sprintf("X is %d! Y is %d!", v.X, v.Y)
// }

// func normal(s string) {
// 	for i := 0; i < 5; i++ {
// 		// time.Sleep(100 * time.Microsecond)
// 		fmt.Println(s)
// 	}
// }

// func goroutine(s string, wg *sync.WaitGroup) {
// 	for i := 0; i < 5; i++ {
// 		// time.Sleep(100 * time.Microsecond)
// 		fmt.Println(s)
// 	}
// 	defer wg.Done()
// }

// func goroutine1(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum
// }

// func producer(ch chan int, i int) {
// 	ch <- i * 2
// }

// func consumer(ch chan int, wg *sync.WaitGroup) {
// 	for i := range ch {
// 		func() {
// 			defer wg.Done()
// 			fmt.Println("process", i*1000)
// 		}()
// 	}
// 	fmt.Println("consumerが終わったよ！！！！！！1")
// }

// func producer(first chan<- int) {
// 	defer close(first)
// 	for i := 0; i < 10; i++ {
// 		first <- i
// 	}
// }

// func multi2(first chan int, second chan<- int) {
// 	defer close(second)
// 	for i := range first {
// 		second <- i * 2
// 	}
// }

// func multi4(second chan int, third chan<- int) {
// 	defer close(third)
// 	for i := range second {
// 		third <- i * 4
// 	}
// }

// func goroutine(s []int, c chan<- int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum
// }

// func goroutine1(ch chan<- string) {
// 	for {
// 		ch <- "packet from 1"
// 		time.Sleep(3 * time.Microsecond)
// 	}
// }

// func goroutine2(ch chan<- int) {
// 	for {
// 		ch <- 45
// 		time.Sleep(1 * time.Microsecond)
// 	}
// }

// type Counter struct {
// 	v   map[string]int
// 	mux sync.Mutex
// }

// func (c *Counter) Inc(key string) {
// 	c.mux.Lock()
// 	defer c.mux.Unlock()
// 	c.v["key"]++
// }

// func (c *Counter) Value(key string) int {
// 	c.mux.Lock()
// 	defer c.mux.Unlock()
// 	return c.v["key"]
// }

func goroutine(s []string, c chan<- string) {
	defer close(c)
	sum := ""
	for _, v := range s {
		sum += v
		c <- sum
	}
}

func main() {
	words := []string{"test1!", "test2!", "test3!", "test4!"}
	c := make(chan string, len(words))
	// c := make(chan string)
	go goroutine(words, c)
	for w := range c {
		fmt.Println(w)
	}

	// c := Counter{v: make(map[string]int)}
	// // c := make(map[string]int)
	// go func() {
	// 	for i := 0; i < 100; i++ {
	// 		c.Inc("key")
	// 	}
	// }()
	// go func() {
	// 	for i := 0; i < 100; i++ {
	// 		c.Inc("key")
	// 	}
	// }()
	// time.Sleep(1 * time.Second)
	// fmt.Println(c.Value("key"))

	// 	tick := time.Tick(100 * time.Microsecond)
	// 	boom := time.After(100 * time.Microsecond)
	// OuterLoop2:
	// 	for {
	// 		select {
	// 		case <-tick:
	// 			fmt.Println("tick")
	// 		case <-boom:
	// 			fmt.Println("BOOM!")
	// 			break OuterLoop2
	// 		}
	// 	}

	// c1 := make(chan string)
	// c2 := make(chan int)
	// go goroutine1(c1)
	// go goroutine2(c2)

	// for {
	// 	select {
	// 	case msg1 := <-c1:
	// 		fmt.Println(msg1)
	// 	case msg2 := <-c2:
	// 		fmt.Println(msg2)
	// 	}
	// }

	// s := []int{1, 2, 3, 4, 5}
	// c := make(chan int)
	// go goroutine(s, c)
	// go goroutine(s, c)
	// x := <-c
	// fmt.Println(x)
	// y := <-c
	// fmt.Println(y)

	// first := make(chan int)
	// second := make(chan int)
	// third := make(chan int)

	// go producer(first)
	// go multi2(first, second)
	// go multi4(second, third)
	// for result := range third {
	// 	fmt.Println(result)
	// }

	// var wg sync.WaitGroup
	// ch := make(chan int)

	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go producer(ch, i)

	// }

	// go consumer(ch, &wg)
	// wg.Wait()
	// close(ch)
	// time.Sleep((2 * time.Second))
	// fmt.Println("Done")

	// x := <-ch
	// fmt.Println(x)

	// s := []int{100, 200}
	// ch := make(chan int, len(s))
	// ch <- s[0]
	// // fmt.Println(ch)
	// // fmt.Println(len(ch))
	// ch <- s[1]
	// // fmt.Println(len(ch))
	// close(ch)

	// for c := range ch {
	// 	fmt.Println(c)
	// }

	// s := []int{1, 2, 3, 4, 5}
	// c := make(chan int)
	// go goroutine1(s, c)
	// // fmt.Println(<-c)
	// x := <-c
	// fmt.Println(x)

	// var wg sync.WaitGroup
	// go goroutine("world", &wg)
	// normal("hello")
	// // time.Sleep(1000 * time.Microsecond)
	// wg.Add(1)
	// wg.Wait()

	// v := Vertex{3, 4}
	// fmt.Println(v)
	// fmt.Println(v.Plus())

	// e := myFunc()
	// fmt.Println(e)

	// e.Error()
	// fmt.Println(e)
	// p := Person{"Kaito", 21}
	// fmt.Println(p)
	// s := p.String()
	// fmt.Println(s)

	// var i interface{} = 10
	// do(i)
	// var s interface{} = "hoge"
	// do(s)

	// var mike Human = &Person{"Mike"}
	// DriveCar(mike)
	// var hachi Dog = Dog{"Hachi"}
	// Naku(hachi)
	// fmt.Println(Naku(hachi))
	// mike.Say()

	// myInt := MyInt(10)
	// fmt.Println(myInt)
	// fmt.Println(myInt.Double())

	// v := New3D(3, 4, 5)
	// v.Scale3D(10)
	// fmt.Println(v.Area3D())

	// v := New(3, 4)
	// v.Scale(10)
	// fmt.Println(v.Area())

	// v := Vertex{3, 4}
	// fmt.Println(v.Area())
	// v.Scale(10)
	// fmt.Println(v.Area())

	// var i int = 100
	// var j int = 200
	// var i_p *int    // pointer(0)
	// var j_p *int    // poitner(0)
	// i_p = &i        // i_p = iのpointer(100)
	// j_p = &j        // j_p = jのpointer(200)
	// i = *i_p + *j_p // 300の実態
	// j_p = i_p       // iのpointer(300)
	// j = *j_p + i    // jのpointer(300) + 300
	// fmt.Println(j)

	// a := 1
	// a = 2 + 3
	// fmt.Println(a)

	// i := 100
	// j := 200
	// var p1 *int
	// var p2 *int
	// p1 = &i // p1 = pointer(100)
	// p2 = &j // p2 = pointer(200)
	// // fmt.Println(p1, p2)
	// i = *p1 + *p2 // 300 = 100 + 200
	// p2 = p1       // pointer(100)
	// j = *p2 + i    // j = 100 + 300
	// fmt.Println(j) // 400➡️答え600!!なぜ？

	// i := 10
	// var p *int // pointer
	// p = &i     // 10
	// var j int  // 0
	// j = *p     //10
	// fmt.Println(j)

	// v1 := Vertex{1, 2, "test"}
	// c1(v1)
	// fmt.Println(v1)

	// v2 := &Vertex{1, 2, "test"}
	// c2(v2)
	// fmt.Println(*v2)

	// var p *Person

	// p = &Person{
	// 	Name: "太郎",
	// 	Age:  20,
	// }
	// fmt.Printf("p :%+v\n", p)
	// fmt.Printf("変数pに格納されているアドレス :%p", p)

	// h := new(int)
	// fmt.Println(h)
	// var p *int = new(int)
	// fmt.Println(p)
	// fmt.Println(*p)
	// *p = 4
	// fmt.Println(*p)

	// var p2 *int
	// fmt.Println(p2)
	// *p2++
	// fmt.Println(p2)
	// var a int = 1
	// a := 1
	// fmt.Println(a)
	// var n int = 100
	// a := one(&n)
	// fmt.Println(a)
	// fmt.Println(n)
	// fmt.Println(&n)

	// l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	// // result1 := 999999999
	// var result1 int
	// for _, num := range l {
	// 	if num < result1 {
	// 		result1 = num
	// 	}
	// }
	// fmt.Println(result1)

	// m := map[string]int{
	// 	"apple":  200,
	// 	"banana": 300,
	// 	"grapes": 150,
	// 	"orange": 80,
	// 	"papaya": 500,
	// 	"kiwi":   90,
	// }
	// result2 := 0
	// // for _, num := range m {
	// // 	result2 += num
	// // }
	// result2 = calc(m)
	// fmt.Println(result2)

	// errors()

	// panik()
	// fmt.Println("fuga")
	// _, err := os.Open("fdfs")
	// if err != nil {
	// 	log.Fatalln("Exit", err)
	// }
	// fmt.Println("piyo")

	// log.Println("log!")
	// fmt.Println(time.Now().Date())
	// m := map[string]int{"apple": 100, "banana": 200}

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// for k := range m {
	// 	fmt.Println(k)
	// }

	// for _, v := range m {
	// 	fmt.Println(v)
	// }

	// i := []string{"a", "b", "c"}
	// for i, v := range i {
	// 	fmt.Println(i, v)
	// }
	// for _, v := range i {
	// 	fmt.Println(v)
	// }
	// for v := range i {
	// 	fmt.Println(v)
	// }

	// for i := 0; i < 1000000; i++ {
	// 	fmt.Println(i)
	// }
	// if i := 0; i == 0 {
	// 	fmt.Printf("t")
	// }
	// Q1
	// f := 1.11
	// fmt.Println(int(f))

	// Q2
	// s := []int{1, 2, 5, 6, 2, 3, 1}
	// fmt.Println(s[2:4])
	// [5,6,2]

	// Q3
	// m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	// fmt.Printf("%T %v", m, m)

	// foo()
	// foo(1)
	// foo(1, 2, 3, 4)

	// s := []int{1, 2, 3}
	// fmt.Println(s)

	// foo(s...)

	// c := ig()
	// fmt.Println(c())
	// fmt.Println(c())
	// fmt.Println(c())
	// fmt.Println(c())

	// c1 := ca(3.14)
	// fmt.Println(c1(2))
	// fmt.Println(c1(3))

	// c2 := ca(3)
	// fmt.Println(c2(2))
	// fmt.Println(c2(3))
	// f := func(x int) { fmt.Println(x) }
	// f(1)
	// f(45)

	// func(x int) {
	// 	fmt.Println(x)
	// }(1)
	// r3 := cal(100, 2)
	// fmt.Println(r3)
	// r1, r2 := add(1, 2)
	// fmt.Println(r1, r2)
	// fmt.Printf("%T", add2("hoge"))

	// b := []byte{72, 73}
	// fmt.Println(b)
	// fmt.Printf(string(b))

	// c := []byte("\nHI")
	// fmt.Println(c)
	// fmt.Println(string(c))
	// m := map[string]int{"apple": 100, "banana": 200}
	// fmt.Println(m)
	// fmt.Println(m["apple"])
	// fmt.Println(m["apple1"])
	// v, ok := m["apple"]
	// fmt.Println(v, ok)

	// m2 := make(map[string]int)
	// m2["pc"] = 5000
	// fmt.Println(m2)
	// fmt.Printf("%T", m["apple"])
	// n := make([]int, 3, 5)
	// fmt.Println(n)
	// fmt.Printf("len=%d cap=%d value=%v", len(n), cap(n), n)
	// n = append(n, 2, 1)
	// fmt.Println("\n", n)
	// fmt.Printf("len=%d cap=%d value=%v", len(n), cap(n), n)
	// n = append(n, 1)
	// fmt.Printf("len=%d cap=%d value=%v", len(n), cap(n), n)

	// c := make([]int, 5)
	// c := make([]int, 0, 5)
	// var c []int = make([]int, 0, 5)
	// for i := 0; i < 5; i++ {
	// 	c = append(c, i)
	// 	fmt.Println(c)
	// }
	// fmt.Println(c)

	// n := []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(n)
	// fmt.Println(n[2:4])
	// fmt.Println(n[:2])
	// fmt.Println(n[1:])
	// fmt.Println(n[:])

	// var board = [][]int{
	// 	[]int{0, 1, 2},
	// 	[]int{3, 4, 5},
	// }
	// fmt.Println(board)
	// // fmt.Println(board[0][1])
	// n = append(n, 5555)
	// fmt.Println(n)
	// board[0] = append(board[0], 5555)
	// fmt.Println(board)

	// var a [2]int
	// a[0] = 100
	// a[1] = 200
	// fmt.Println(a)

	// var b [3]int = [3]int{1, 2, 3}
	// fmt.Println(b)

	// var c []int = []int{1, 2}
	// c = append(c, 4)
	// fmt.Println(c)
	// b = append(b, 4)

	// var x int = 1
	// xx := float64(x)
	// fmt.Printf("%T, %v %f\n", xx, xx, xx)

	// var y float64 = 1.2
	// yy := int(y)
	// fmt.Printf("%T, %v %d\n", yy, yy, yy)

	// var s string = "{"
	// i, errors := strconv.Atoi((s))
	// fmt.Println(errors)
	// fmt.Printf("%T, %v", i, i)

	// t, f := true, false
	// fmt.Printf("%t %t", t, f)
	// const a = 1
	// bazz()
	// s := "hhhi"
	// p := strings.Replace(s, "h", "i", 2)
	// fmt.Println(p)
	// fmt.Println(s)
	// fmt.Println(strings.Contains(s, "i"))
	// fmt.Println(s)
	// fmt.Println(`hoge
	// fuga`)

	// fmt.Println(1+1, 1+1)
	// fmt.Println(1 + 1)
	// fmt.Println(1)
	// fmt.Println("Hello world!", "TEST TEST")
}

func init() {
	fmt.Println("-------------------------- init! --------------------------")
}
