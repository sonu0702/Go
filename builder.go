package  main

import (
	"fmt"
)

type vehicle struct {
	wheel int
	color string
	gas int
}

type vehicleBuilder interface {
	setWheel(int) vehicleBuilder
	setColor(string) vehicleBuilder
	maxGas(int) vehicleBuilder
} 

type carBuilder struct {
	v vehicle
}
 

func (c* carBuilder) setWheel(num int) vehicleBuilder{
	c.v.wheel = num
	return c
}

func (c * carBuilder) setColor(color string) vehicleBuilder{
	c.v.color = color
	return c
}

func (c* carBuilder) maxGas(gas int) vehicleBuilder {
	c.v.gas = gas
	return c
}

type bikeBuilder struct {
	v vehicle
}

func (c* bikeBuilder) setWheel(num int) vehicleBuilder{
	c.v.wheel = num
	return c
}

func (c * bikeBuilder) setColor(color string) vehicleBuilder{
	c.v.color = color
	return c
}

func (c* bikeBuilder) maxGas(gas int) vehicleBuilder {
	c.v.gas = gas
	return c
}

//ManufactureManager build what ever has vehicleBuilder
type ManufactureManager struct {
	builder vehicleBuilder
}

func (m* ManufactureManager) setBuilder(b vehicleBuilder){
	m.builder = b
}

func (m* ManufactureManager) consttruct(wheel int,color string,gas int  ){
	m.builder.setWheel(wheel).setColor(color).maxGas(gas)
}


func main(){
	var buildMan ManufactureManager
	var car *carBuilder
	car = &carBuilder{}
	buildMan.setBuilder(car)
	buildMan.consttruct(4,"Pink",120)

	var bike *bikeBuilder
	bike = &bikeBuilder{}
	buildMan.setBuilder(bike)
	buildMan.consttruct(2,"Pink",220)

	fmt.Println(car)
	fmt.Println(bike)
}

