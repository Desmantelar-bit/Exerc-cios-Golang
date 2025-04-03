package main

import "fmt"

func main() {
	var contagemRegressiva = [] int{90, 80, 70, 60, 50}
	contagemRegressiva = append(contagemRegressiva, 40, 30, 20)
	fmt.Println(contagemRegressiva)
	fmt.Println(len(contagemRegressiva), cap(contagemRegressiva))

}
