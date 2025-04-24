package main

import "fmt"

func main() {
	
	produtosCantina := map[string]int{
		"Coxinha": 10,
		"Pão de Queijo": 15,
		"Refrigerante":  20,
	}
	 for produto, quantias := range produtosCantina {
		fmt.Printf("Produto: %s; Quantidade: %d\n", produto, quantias)
	}
	produtosCantina["Coxinha"] -= 2
	produtosCantina["Pão de Queijo"] -= 1

	fmt.Println("Após a venda:")
	for produto, quantias := range produtosCantina {
		fmt.Printf("Produto: %s; Quantidade: %d\n", produto, quantias)
	}
}

