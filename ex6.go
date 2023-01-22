package main

import "fmt"

func main(){

	s := "Çéou oei?"

	// scrolls through the string by characteristic
	for _, v := range s {
		fmt.Printf("%b - %T - %#U - %#X\n", v, v, v, v)
	}

	fmt.Println("")

	// scrolls through the string by byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%b - %T - %#U - %#X", s[i], s[i], s[i], s[i])
	}

}