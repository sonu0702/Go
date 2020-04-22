## Example 1
```go
type geometry interface{
	area() float64
}

type rec struct {
	width float64
	height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64{
	return c.radius * math.Pi
}

func shapes(g geometry) float64{
	fmt.Println(g);
	return g.area()
}
func main(){
	cir := circle{2.0}
	value := shapes(cir)
	fmt.Println(value)	
}
```
## Example 2

```go

package main

import (
	"fmt"
)

type writer interface {
	write()
}

type reader interface{
	read()
}

type writeread interface {
	reader
	writer
}

type data struct {
	value string
}

func (d data) read(){
	fmt.Println("I am reading")
}

func (d data) write(){
	fmt.Println("I am writing")
}

func operation(wr writeread){
	wr.read()
	wr.write()
}


func main(){
	fmt.Println("data fetched")
	dd := data{"Nothing"}
	operation(dd)
}
```

## Example 3

```go

package main

import (
	"fmt"
)

type objStr struct {}


type noStr struct {}

type generator interface {
	foo()
}

func (o objStr) foo(){
	fmt.Println(o)
}

func simple() generator {
	return &objStr{}
}

func main(){
	fmt.Println("data fetched")
	ov := simple()
	fmt.Println(ov)
}

```

## Example 4

```go
package main

import (
	"fmt"
)

type myValuer interface{
	myValue()
}

type gitObject struct{
	value int
}

func (g *gitObject) myValue(){
	fmt.Printf("my value by pointer %d",g.value)
}

func retP() myValuer{
	return &gitObject{2}
}

func main(){
	fmt.Println("data fetched")
	wp := retP()
	wp.myValue()
}
```

