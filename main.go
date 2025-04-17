package main

import ("fmt")

func dadosPessoais(idade int, nome string,) (int, string) {
	
	if idade < 0 || idade > 120 {
		fmt.Println("Dados inválidos.")
	} else if idade < 18 || idade > 0{
		fmt.Println("Você é menor de idade.")
	} else if idade >= 18 && idade < 60 {
		fmt.Println("Você é maior de idade.")
	} else if idade >= 60 {
		fmt.Println("Você é idoso.")	
	}
	
	return idade, nome
}

func main() {
	var idade int
	var nome string

	var letras = []string{nome}
	fmt.Println("Qual o seu nome?")
	fmt.Scan(&nome)
	fmt.Println("Qual a sua idade?")
	fmt.Scan(&idade)
	fmt.Println("Nome:", nome)
	fmt.Println("Letras do nome:", len(letras))
	fmt.Println("Idade:", idade)
	dadosPessoais(idade, nome)
}
