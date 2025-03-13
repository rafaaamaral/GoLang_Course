package main

import "fmt"

func main() {
	// fmt.Println("Loops")
	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println(i)
	// }

	// for j := 0; j < 10; j++ {
	// 	fmt.Println(j)
	// }

	nomes := [3]string{"João", "Davi", "Lucas"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

}
