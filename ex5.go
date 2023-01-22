package main

import(
	"fmt"
	"runtime"
) 

func main(){
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	overflow()
}

func overflow(){
	var x uint16
	x = 65535
	fmt.Println(x)
	x ++ 
	fmt.Println(x)
	x ++ 
	fmt.Println(x)

}