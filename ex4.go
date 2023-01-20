package main

import "fmt"

type meu_tipo int
var x meu_tipo

func main(){
	fmt.Printf("%v %T", x, x)
	x = 42
	fmt.Println(x)
	y := int(x)

	fmt.Printf("%v %T", y, y)
}
