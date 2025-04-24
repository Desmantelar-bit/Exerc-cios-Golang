package main

import "fmt"

func main() {
	capitais := map[string]string{
		"SP": "São Paulo",
		"RJ": "Rio de Janeiro",
		"MG": "Minas Gerais",
		"BH": "Bahia",
	}

	capitais["PR"] = "Paraná"

	for sigla, capital := range capitais {
		fmt.Printf("A capital do estado %s é %s\n", sigla, capital)
	}
	delete(capitais, "BH")
	fmt.Println("Após a remoção de BH:")
	for sigla, capital := range capitais {
		fmt.Printf("A capital do estado %s é %s\n", sigla, capital)
	}

	matiz := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for _, linha := range matiz {
		fmt.Println(linha)
	}
	
}

