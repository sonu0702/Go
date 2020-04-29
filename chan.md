## Example
```go
package main

import (
	"fmt"
)

func worker(done chan bool){
	fmt.Println("I am doing work...")
	done <- true
}


func main(){
	done := make(chan bool)
	go worker(done)
	<- done
}
```

## Example
```go
package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
	}
	fmt.Println(time.Now())
}
```
## Example

```go
package main

import (
	"math/rand"
	"errors"
	"fmt"
	"time"
)

//Result d
type Result struct {
	Title string
	URL   string
}

//SearchFunc d
type SearchFunc func(query string) Result

//FakeSearch is to FakeSearch
func FakeSearch(kind, title, url string) SearchFunc {

	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{
			Title: fmt.Sprintf("%s(%q) : %s", kind, query, title),
			URL:   url,
		}
	}
}

var (
	web   = FakeSearch("web", "This is Web", "https://fake-search-url/")
	image = FakeSearch("image", "This is image", "https://fake-search-url")
	video = FakeSearch("video", "This is Video", "https://fake-search-url")
)

//Search gives result
func Search(query string) ([]Result, error) {

	res := []Result{
		web("where is india"),
		image("my mom"),
		video("hello done"),
	}
	return res, nil

}

//SearchParallel for result
func SearchParallel(query string) ([]Result, error) {
	c := make(chan Result, 3)
	go func() { c <- web("golang") }()
	go func() { c <- image("golang") }()
	go func() { c <- video("golang") }()
	return []Result{<-c, <-c, <-c}, nil
}

//SearchTimeout with result
func SearchTimeout(query string, timeout time.Duration) ([]Result, error) {
	timer := time.After(timeout)
	c := make(chan Result, 3)
	go func() { c <- web("golang") }()
	go func() { c <- image("golang") }()
	go func() { c <- video("golang") }()

	var results []Result
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)

		case <-timer:
			return results, errors.New("timedout")
		}
	}
	return results,nil
}
func main() {
	start := time.Now()
	fmt.Println(SearchTimeout("Golang", 100*time.Millisecond))
	// fmt.Println(SearchParallel("golang"))
	fmt.Println(time.Since(start))
}

```
