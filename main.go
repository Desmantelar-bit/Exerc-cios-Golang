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

	fmt.Println("Quantas coxinhas foram vendidas?")
	var coxinhasVendidas int
	fmt.Scan(&coxinhasVendidas)
	fmt.Println("Quantos pães de queijo foram vendidos?")
	var paesDeQueijoVendidos int
	fmt.Scan(&paesDeQueijoVendidos)
	fmt.Println("Quantos refrigerantes foram vendidos?")
	var refrigerantesVendidos int
	fmt.Scan(&refrigerantesVendidos)
	
	produtosCantina["Coxinha"] -= coxinhasVendidas
	produtosCantina["Refrigerante"] -= refrigerantesVendidos
	produtosCantina["Pão de Queijo"] -= paesDeQueijoVendidos

	fmt.Println("Após a venda:")
	for produto, quantias := range produtosCantina {
		fmt.Printf("Produto: %s; Quantidade: %d\n", produto, quantias)
	}
}

