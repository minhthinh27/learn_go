package useful_closure

import (
	"fmt"
	"net/http"
	"sort"
	"time"
)

// https://www.calhoun.io/5-useful-ways-to-use-closures-in-go/
func Run() {

	// gen := makeFibGen()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(gen())
	// }

	// http.HandleFunc("/hello", timed(hello))
	// http.ListenAndServe(":3000", nil)

	// db := NewDatabase("localhost:5432")
	// http.HandleFunc("/hello", hello1(db))
	// http.ListenAndServe(":3000", nil)

	// Binary searching with the sort package
	numbers := []int{1, 11, -5, 8, 2, 0, 12}
	sort.Ints(numbers)
	fmt.Println("Sorted:", numbers)

	index := sort.Search(len(numbers), func(i int) bool {
		return numbers[i] >= 7
	})
	fmt.Println("The first number >= 7 is at index:", index)
	fmt.Println("The first number >= 7 is:", numbers[index])

	// Deferring work

}

// Isolating data
func makeFibGen() func() int {
	f1 := 0
	f2 := 1

	return func() int {
		f2, f1 = (f1 + f2), f2
		return f1
	}
}

// Wrapping functions and creating middleware
func timed(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		end := time.Now()
		fmt.Println("The request took", end.Sub(start))
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}

// Accessing data that typically isn’t available

type Database struct {
	Url string
}

func NewDatabase(url string) Database {
	return Database{url}
}

func hello1(db Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, db.Url)
	}
}
