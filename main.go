package main

import "fmt"

func main() {
	
	alunoIdade := make(map[string]int)
	alunoIdade["Bruno"] = 15
	alunoIdade["Otávio"] = 16
	alunoIdade["Fabiano"] = 40
	alunoIdade["Isabela"] = 15
	fmt.Println("Idade do Bruno", alunoIdade["Bruno"])

	notasAluno := map[string]float64{
		"Bruno":   7.5,
		"Otávio":  8.0,
		"Fabiano": 9.0,
		"Isabela": 10.0,
	}

for nome, nota := range notasAluno {
	fmt.Printf("%s tem nota %.2f\n", nome, nota)}
}
