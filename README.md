## Notes
1.  factored imports 
```go
import (
"name"

"newname"
)
```

2.  exported function are capitiled

```go
math.pi ✗

math.Pi ✓
```
  3.  Types after variable `func add( x int , y int) int {}`
  
  4.  have consecutive named function parameteres omit the all types but for the last one.
  ```go
  func add(x, y int) int {}
  
  ```
  5. func can return any number of argumanets
    ``` go
      func get(x,y string)(string ,string){ return x ,y}
      ...
      x,y := get("one","two")
    ```
  6. Naked return - func definition for return can be named so these are treated as variables , writing just a return
    , returns these named varible values.
      ```go 
      func get(x,y string)(x string , y string){x ,y}
      ...
      x,y := get("one","two")
     ```
     
  7. `var python, java bool` var declares variables at function or package level 
  
  8. `var c, python, java = true, false, "no!"` var initialization  ,if type is omitted var type is infered from values.
  
  9. inside a func , var can be omitted `k := 6`
  
  10 . factored declarration
  ```go
  var (
	  ToBe   bool       = false
	  MaxInt uint64     = 1<<64 - 1
	  z      complex128 = cmplx.Sqrt(-5 + 12i)
  )
  ```
  
  11. The ### Zero type values , var decared without initial values are given their zero values
  ```go
  var (
    i int 
    x bool
    y string
  )
  0   - int
  false - bool
  ""   - string
  ```
  
  12. The expression T(v) converts the value v to the type T.
  
  13. `const` is same as var but can't use := with const

## Flow Controls
1. `for` only one loop construct , 
```go for i :=0; i<10; i++ {}```
2. init and post statements are optional, avod ;; for while loop :smile: `while` loop.
3. `if` are same as C , surprice1 is 
```go
  if v := math.Pow(x, n); v < lim {
  
		return v
    
	}  
  ```
  here `v` scope is under if. surprice2 v is also available in any of the `else` blocks

4. `switch` statements run only one case which matches  , i.e do not need to provide `break` like python,c etc
5. `defer` statements defers the execution of function until the surrounding function return
6. Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order
```go
for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
  //prints in reverse order
 ```
## Pointers

1. The type *T is a pointer to a T value. Its zero value is nil.
2. The & operator generates a pointer to its operand.
3. The * operator denotes the pointer's underlying value.
4. *p = 21         // set i through the pointer p
5. This is known as "dereferencing" or "indirecting".
6. Unlike C, Go has no pointer arithmetic

```go
package main

import "fmt"

func main() {
	var p *int
	
	i := 0
	
	p = &i
	
	*p = 40
	
	fmt.Println(i)
  
}
```

## Arrays
1.  The type [n]T is an array of n values of type T.
```go
var a[10]int

primes := [6]int{2, 3, 5, 7, 11, 13}

```
2.  An array's length is part of its type, so arrays cannot be resized. 

## Slices
1.  The type []T is a slice with elements of type T.
2.  Flexible length.
3.  Slices do not store data , they work on underlying array data.
4.  Changing value in slice changes values in original array as well as the other slices reffering to that array.
5.  The length of a slice is the number of elements it contains.
6.  The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
7.  The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
8.  The zero value of a slice is nil.
9.  Can create dynamic Slice using make
```go
import "fmt"

func main() {

	a := make([]int, 5)  // len 5 , cap 5
   
	printSlice("a", a)

	b := make([]int, 0, 5) // len 0 , cap 5
  
	printSlice("b", b)

	c := b[:2]
  
	printSlice("c", c)

	d := c[2:5]
  
	printSlice("d", d)
  
}

func printSlice(s string, x []int) {

	fmt.Printf("%s len=%d cap=%d %v\n",
  
		s, len(x), cap(x), x)
    
}
```go
10. Slices can contain any type, including other slices
11. `s = append(s, 2, 3, 4)`  appends values to slice.
12. The range form of the for loop iterates over a slice or map.
```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

	for i, v := range pow {  // can use for _, v := range pow { } to skip index
  
		fmt.Printf("2**%d = %d\n", i, v)
    
	}
  
}
```
## Maps
1.  The zero value of a map is nil. A nil map has no keys, nor can keys be added.
2.  The make function returns a map of the given type, initialized and ready for use.
```go
package main

import "fmt"

type vertex struct{

  lat,long int64

}

var m map[string]vertex

func main(){
	
  m = make(map[string]vertex)
	
  fmt.Println(m)
	
  m["new"] = vertex{23 , 56}
	
  fmt.Println(m["new"])
  delete(m, "new")

}
```

3. `v, ok := m["Answer"]` v is value of "Answer" and ok is false or true according to availibility of key
 
```go
package main

import (
	"golang.org/x/tour/wc"
		"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	result := map[string]int{"x": 1}
	str := strings.Fields(s)
	fmt.Println(str)
	for _,v := range str {
		res , ok := result[v]
		if ok == true {
			result[v] = res + 1
		} else {
			result[v] =  1
		}
		
		fmt.Println(result)
	}
	delete(result , "x")
	return result
	
}

func main() {
	wc.Test(WordCount)
}

```go
## Function values
1. Just linke javascript.

##  Function closures
1.  Go functions may be closures. 
2.  A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
 
 ```go
package main

import "fmt"

func adder() func(int) int {

sum := 0
	
  return func(x int) int {
		
    sum += x
		
    return sum
	
  }

}


func main() {
	
  pos, neg := adder(), adder()
	
  for i := 0; i < 10; i++ {
		
    fmt.Println(
			
      pos(i),
			
      neg(-2*i),
		
    )
	
  }

}

 ```

## Fib in Closure
```go
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n := 0
	prev := 0
	next := 1
	c 	 := prev + next
	return func() int {
		var ret int
		switch {
			case n == 0 :
				n++
				ret = 0
			
			case n == 1 :
				n++
				ret = 1
			default :
				ret = c
				prev = next
				next = c
				c = prev + next
		}
		return ret
	}
}	

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```

## Methods and Interfaces

## Goroutines

1.  A goroutine is a lightweight thread managed by the Go runtime.
2.  Goroutines run in the same address space, so access to shared memory must be synchronized.

## Channels
1.  Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
```
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```
```go

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
```
##  Buffered Channels
Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:
```go
ch := make(chan int, 100)
```



 
 
 










