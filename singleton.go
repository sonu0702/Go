 package  main

import (
	"sync"
	"fmt"
	"math/rand"
)

var once sync.Once

//Singleton is single
type Singleton struct{
	val int
}

var instance *Singleton

func getInstnce() *Singleton{
	v :=  rand.Intn(100)
	fmt.Println(v)
	once.Do(func (){
		instance = &Singleton{
			val :v,
		}
	})
	return instance
}

func main(){
	fmt.Println(getInstnce())
	fmt.Println(getInstnce())
	fmt.Println(getInstnce())
}
