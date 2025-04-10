package main

import "fmt"

func main() {
	
	numero := [10]float32{}
	fmt.Println("Digite 1 numero")
	fmt.Scan(&numero[0])
	fmt.Println("Digite o segundo numero")
	fmt.Scan(&numero[1])
	fmt.Println("Digite o terceiro numero")
	fmt.Scan(&numero[2])
	fmt.Println("Digite o quarta numero")
	fmt.Scan(&numero[3])
	fmt.Println("Digite o quinto numero")
	fmt.Scan(&numero[4])
	fmt.Println("Agora somaremos os 5 numeros")
	fmt.Println("A soma dos 5 numeros é: ", numero[0] + numero[1] + numero[2] + numero[3] + numero[4])
	fmt.Println("A media dos 5 numeros é: ", (numero[0] + numero[1] + numero[2] + numero[3] + numero[4]) / 5)
}
