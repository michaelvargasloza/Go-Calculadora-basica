package main

import "fmt"

func main() {
	var num1, num2 float64

	fmt.Print("Ingresa el primer número: ")
	fmt.Scan(&num1)

	fmt.Print("Ingresa el segundo número: ")
	fmt.Scan(&num2)

	resultado := num1 + num2

	fmt.Printf("La suma de %.2f y %.2f es %.2f\n", num1, num2, resultado)
}
