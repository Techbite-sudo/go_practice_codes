package main

import (
	"fmt"
)

type Shape interface{
	area() float32
}
type Square struct{
	length float32
}
type Rectangle struct{
	length,width float32
}

func (s Square) area() float32{
	return s.length * s.length
}

func(r Rectangle) area() float32{
	return r.length *r.width
}

func getArea(s Shape) float32{
	return s.area()
}
func main(){
  square1:=Square{length: 4}
  rect1:=Rectangle{length: 5, width: 8}
  fmt.Printf("Area is %f\n",getArea(square1) )
  fmt.Printf("Area is %f\n",getArea(rect1) )
}