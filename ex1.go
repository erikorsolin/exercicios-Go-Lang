package main

import "fmt"

type meu_tipo int

var b meu_tipo

func main() {
	n1 := 20
	n2 := 20
	fmt.Println("rodando")
	fmt.Printf("%T\n", b)
	soma(n1, n2)
	multiplica(n1, n2)
	divide(n1, n2)
	imprime_text()

}

func soma(x int, y int) {

	fmt.Println("a soma é:", x+y)

}

func multiplica(x int, y int) {
	fmt.Println("a multiplicação é:", x*y)

}

func imprime_text() {
	x := "abc"
	y := "def"
	z := "ghi"
	w := fmt.Sprint(x, y, z)
	fmt.Println(w)
}

func divide(x int, y int) {
	fmt.Println("A divisão é:", x/y)
}


