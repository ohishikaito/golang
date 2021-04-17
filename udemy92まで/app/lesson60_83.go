package main

import (
	"fmt"
)

// type Person struct {
// 	Name      string   `json:-`
// 	Age       int      `json:age,omitempty`
// 	Nicknames []string `json:nicknames`
// }

// func (p *Person) UnmarshalJSON(b []byte) error {
// 	type Person2 struct {
// 		Name string
// 	}
// 	var p2 Person2
// 	err := json.Unmarshal(b, &p2)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	p.Name = p2.Name + "!"
// 	return err
// }

// var DB = map[string]string{
// 	"User1Key": "User1Secret",
// 	"User2Key": "User2Secret",
// }

// func Server(apiKey, sign string, data []byte) {
// apiSecret := DB[apiKey]
// h := hmac.New(sha256.New, []byte(apiSecret))
// h.Write(data)
// expectedHMAC := hex.EncodeToString(h.Sum(nil))
// fmt.Println(sign == expectedHMAC)
// fmt.Println(apiSecret)
// }

// func increment(i *int) {
// 	*i++
// 	fmt.Println(*i)
// }

// func inc2(i2 *int) {
// 	*i2++
// 	fmt.Println(*i2)
// }

// var s *semaphore.Weighted = semaphore.NewWeighted(1)

// func longProcess(ctx context.Context) {
// isAcquire := s.TryAcquire(1)
// if !isAcquire {
// 	fmt.Println("no")
// 	return
// }
// if err := s.Acquire(ctx, 1); err != nil {
// 	fmt.Println(err)
// }
// defer s.Release(1)
// fmt.Println("Wait...")
// time.Sleep(1 * time.Second)
// fmt.Println("Done")
// }

// type ConfigList struct {
// 	Port      int
// 	DbName    string
// 	SQLDriver string
// }

// var config ConfigList

// type JsonRPC2 struct {
// 	Version string      `json:"jsonrpc"`
// 	Method  string      `json:"method"`
// 	Params  interface{} `json:"params"`
// 	Result  interface{} `json:"result,omitempty"`
// 	Id      *int        `json:"id,omitempty"`
// }
// type SubscribeParams struct {
// 	Channel string `json:"channel"`
// }

func main() {
	// u := url.URL{Scheme: "wss", Host: "ws.lightstream.bitflyer.com", Path: "/json-rpc"}
	// log.Printf("connecting to %s", u.String())

	// c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	// if err != nil {
	// 	log.Fatal("dial:", err)
	// }
	// defer c.Close()

	// if err := c.WriteJSON(&JsonRPC2{Version: "2.0", Method: "subscribe", Params: &SubscribeParams{"lightning_ticker_BTC_JPY"}}); err != nil {
	// 	log.Fatal("subscribe:", err)
	// 	return
	// }

	// for {
	// 	message := new(JsonRPC2)
	// 	if err := c.ReadJSON(message); err != nil {
	// 		log.Println("read:", err)
	// 		return
	// 	}

	// 	if message.Method == "channelMessage" {
	// 		log.Println(message.Params)
	// 	}
	// }

	// cfg, _ := ini.Load("config.ini")
	// config = ConfigList{
	// 	Port:      cfg.Section("web").Key("port").MustInt(),
	// 	DbName:    cfg.Section("web").Key("port").MustString("example.sql"),
	// 	SQLDriver: cfg.Section("db").Key("driver").String(),
	// }

	// fmt.Printf("%T %v\n", config.Port, config.Port)
	// fmt.Printf("%T %v\n", config.DbName, config.DbName)
	// fmt.Printf("%T %v\n", config.SQLDriver, config.SQLDriver)

	// ctx := context.TODO()
	// go longProcess(ctx)
	// go longProcess(ctx)
	// go longProcess(ctx)
	// time.Sleep(2 * time.Second)
	// go longProcess(ctx)
	// time.Sleep(5 * time.Second)

	// i := 1
	// i2 := 10
	// go increment(&i)
	// go inc2(&i2)
	// go increment(&i)
	// go inc2(&i2)
	// go increment(&i)
	// go inc2(&i2)

	// time.Sleep(5 * time.Second)
	// const apiKey = "User2Key"
	// const apiSecret = "User2Secret"

	// data := []byte("data")
	// // c := sha256.New
	// h := hmac.New(sha256.New, []byte(apiSecret))
	// h.Write(data)
	// sign := hex.EncodeToString(h.Sum(nil))

	// fmt.Println(sign)

	// Server(apiKey, sign, data)

	// b := []byte(`{"name":"mike","age":0,"nicknames":["a","b","c"]}`)
	// var p Person
	// if err := json.Unmarshal(b, &p); err != nil {
	// 	fmt.Println(err)
	// }
	// // fmt.Println(p)
	// fmt.Println(p.Name, p.Age, p.Nicknames)

	// v, _ := json.Marshal(p)
	// fmt.Println(string(v))

	// p.UnmarshalJSON(b)

	// base, _ := url.Parse("http://example.com")
	// fmt.Println(base)

	// ref, _ := url.Parse("/test?a=1&b=2")
	// ep := base.ResolveReference(ref).String()
	// fmt.Println(ep)

	// // req, _ := http.NewRequest("GET", ep, nil)
	// params := bytes.NewBuffer([]byte("password"))
	// req, _ := http.NewRequest("POST", ep, params)
	// req.Header.Add("If-None-Match", "hoge")
	// q := req.URL.Query()
	// q.Add("c", "333")
	// req.URL.RawQuery = q.Encode()

	// fmt.Println(req)

	// var client *http.Client = &http.Client{}
	// resp, _ := client.Do(req)
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	// resp, _ := http.Get("http://example.com")
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(resp)
	// fmt.Println(body)
	// r := bytes.NewBuffer(([]byte("123")))
	// // r := bytes.newBuffer([]byte("abc"))
	// content, what := ioutil.ReadAll(r)
	// fmt.Println(string(content), what)

	// content, err := ioutil.ReadFile("app/main.go")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(content)
	// fmt.Println(string(content))

	// if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil {
	// 	log.Fatalln(err)
	// }

	// const a = iota
	// const b = iota
	// const c = iota
	// fmt.Println(a, b, c)

	// i := []int{5, 3, 2, 8, 7}
	// s := []string{"d", "a", "f"}
	// p := []struct {
	// 	Name string
	// 	Age  int
	// }{
	// 	{"Nancy", 20},
	// 	{"Vera", 40},
	// 	{"Mike", 30},
	// 	{"Bob", 50},
	// }
	// fmt.Println(i, s, p)
	// sort.Ints(i)
	// sort.Strings(s)
	// // sort.Slice(p, func())
	// sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	// // sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })
	// fmt.Println(i, s, p)

	// match, v := regexp.MatchString("a([a-z]+)e", "apple")
	// fmt.Println(match, v)

	// r := regexp.MustCompile("a([a-z]+)e")
	// ms := r.MatchString("apple")
	// fmt.Println(ms)

	// r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	// fs := r2.FindString(("/view/test"))
	// fmt.Println(fs)

	// t := time.Now()
	// fmt.Println(t)
	// fmt.Println(t.Format(time.RFC3339))
	// fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	// fmt.Println(t.Date())

	// spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	// fmt.Print(spy.CSV())
	// rsi2 := talib.Rsi(spy.Close, 2)
	// fmt.Println(rsi2)

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
