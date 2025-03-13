package main

import "fmt"

func main() {
	usuarios := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuarios)

	usuarios2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Silva",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuarios2)

	delete(usuarios2, "nome")
	fmt.Println(usuarios2)

	usuarios2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}

	fmt.Println(usuarios2)
}
