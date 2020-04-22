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
