package main

import ("fmt" 
	 "math")

func main() {

	C1 := circle{10}

	fmt.Println ("circumfarence of circle is =" circumfarence())
}

//type Shape interface {
//
//	area() float64
//}

type circle interface{
	redius float64	
}

func (C1 *circle) circumfarence() float64{

	return (2 * math.pi * redius)
}
	
}