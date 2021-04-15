package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// type User struct {
// 	ID   int
// 	Name string
// }

// type Person struct {
// 	Id   int
// 	Name string
// 	Age  int
// }

// var testCases = []struct {
// 	in   string
// 	want map[string]int
// }{
// 	// {"I am learning Go!", map[string]int{
// 	// 	"I": 1, "am": 1, "learning": 1, "Go!": 1,
// 	// }},
// 	// {"The quick brown fox jumped over the lazy dog.", map[string]int{
// 	// 	"The": 1, "quick": 1, "brown": 1, "fox": 1, "jumped": 1,
// 	// 	"over": 1, "the": 1, "lazy": 1, "dog.": 1,
// 	// }},
// 	{"I ate a donut. Then I ate another donut.", map[string]int{
// 		"I": 2, "ate": 2, "a": 1, "donut.": 2, "Then": 1, "another": 1,
// 	}},
// 	// {"A man a plan a canal panama.", map[string]int{
// 	// 	"A": 1, "man": 1, "a": 2, "plan": 1, "canal": 1, "panama.": 1,
// 	// }},
// }

// func Test(f func(string) map[string]int) {
// 	ok := true
// 	for _, c := range testCases {
// 		got := f(c.in)
// 		if len(c.want) != len(got) {
// 			ok = false
// 		} else {
// 			// fmt.Println(c.want, got)
// 			for k, _ := range c.want {
// 				// fmt.Println(k, v)
// 				// fmt.Println(c.want[k], got[k])
// 				if c.want[k] != got[k] {
// 					ok = false
// 				}
// 			}
// 		}
// 		if !ok {
// 			fmt.Printf("FAIL\n f(%q) =\n  %#v\n want:\n  %#v",
// 				c.in, got, c.want)
// 			break
// 		}
// 		fmt.Printf("PASS\n f(%q) = \n  %#v\n", c.in, got)
// 	}
// }

// func WordCount(s string) map[string]int {
// 	sf := strings.Fields(s)
// 	fmt.Println(sf, "【sf】")
// 	num := len(sf)
// 	ret := make(map[string]int)
// 	for i := 0; i < num; i++ {
// 		fmt.Println(sf[i])
// 		(ret[sf[i]]) += 1

// 		ret["hoge"] += 1

// 		// map配列だから、キーに対して勝手にvalueが入るって感じなのかな

// 		// (ret["I": 2])
// 		fmt.Println(ret)
// 	}
// 	fmt.Println(ret["I"])
// 	return ret
// }

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

	// p1 := &Page{Title: "test", Body: []byte("sample")}
	// p1.save()

	// // fmt.Println(string(p1.Title))
	// p2, _ := loadPage(p1.Title)
	// // fmt.Println(p1, p2)
	// fmt.Println(string(p2.Body))
	// Test(WordCount)

	// DbConnection, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3308)/go_database")
	// defer DbConnection.Close()
	// cmd := `CREATE TABLE IF NOT EXISTS persons(
	// 			name TEXT,
	// 			age  INT)`
	// _, err = DbConnection.Exec(cmd)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmd = "INSERT INTO persons (name, age) VALUES (?, ?)"
	// _, err = DbConnection.Exec(cmd, "Mike", 24)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmd = "UPDATE persons SET age = ? WHERE NAME = ?"
	// _, err = DbConnection.Exec(cmd, 25, "Mike")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmd = "SELECT * FROM persons"
	// rows, err := DbConnection.Query(cmd)
	// defer rows.Close()
	// var pp []Person
	// for rows.Next() {
	// 	var p Person
	// 	err := rows.Scan(&p.Id, &p.Name, &p.Age)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	pp = append(pp, p)
	// }
	// for _, p := range pp {
	// 	fmt.Println(p.Name, p.Age)
	// }

	// for rows.Next() {
	// 	var user User
	// 	err := rows.Scan(&user.ID, &user.Name)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	fmt.Println(user.ID, user.Name)
	// }

	// err = rows.Err()
	// if err != nil {
	// 	panic(err.Error())
	// }

	// cmd = "SELECT * FROM users WHERE NAME = ?"
	// row := DbConnection.QueryRow(cmd, "太郎")
	// var user User
	// err = row.Scan(&user.ID, &user.Name)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No row")
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }
	// fmt.Println(user.ID, user.Name)

	// cmd = "DELETE FROM persons WHERE age = 24"
	// _, err = DbConnection.Exec(cmd)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}

func init() {
	fmt.Println("-------------------------- init! --------------------------")
}
