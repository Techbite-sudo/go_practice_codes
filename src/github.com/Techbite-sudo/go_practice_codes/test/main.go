package main

import (
	"fmt"
	"math"
)

type Shape interface{
	area() float64
}
type Circle struct{
	radius float64
}
type Rectangle struct{
	length,width float64
}
func (c Circle) area() float64{
	return math.Pi * c.radius * c.radius
}
func (r Rectangle) area() float64{
	return r.length * r.width
}
func getArea(s Shape) float64{
	return s.area() 
}

func main(){
	circle :=Circle{radius:7}
	rectangle :=Rectangle{length: 5,width: 7}

	fmt.Println("Circle Area = ",getArea(circle))
	fmt.Println("Rectangle Area = ",getArea(rectangle))
}