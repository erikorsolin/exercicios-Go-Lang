package main

import "fmt"

const x float64 = 55.5  // tipado
const y int = 10 // tipado
const z = "wewe"  // nao tipado

func main(){
	fmt.Printf("%v %T\n%v %T\n%v %T\n", x, x, y, y, z, z)
}