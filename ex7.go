package main

import "fmt"

const y = 10.0
var b float64
const(
	_ = iota * 10
	a
	w
	c
	d
	e

)

func main(){
	b = y
	fmt.Printf("%v %T  ", b, b)
	iotas()
}


func iotas(){
	fmt.Println(a, w, c, d, e)

}