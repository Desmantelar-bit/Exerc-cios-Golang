package main

import "fmt"

func main() {
	/*var contagemRegressiva = [] int{90, 80, 70, 60, 50}
	contagemRegressiva = append(contagemRegressiva, 40, 30, 20)
	fmt.Println(contagemRegressiva)
	fmt.Println(len(contagemRegressiva), cap(contagemRegressiva))
*/
	var nomes = [] string{"Lucas", "João", "Sabrina", "Maria", "Ana"}
	fmt.Println(nomes [0:2])
	fmt.Println(nomes [3:5])
	rangeNomedoMeio := nomes [2]
	fmt.Println(rangeNomedoMeio)
}
