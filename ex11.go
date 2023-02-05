package main

import "fmt"

var b = 100

func main(){
	fmt.Printf("%v %X %b\n", b, b, b)
	w := b << 2
	fmt.Printf("%v %X %b", w, w, w)
}