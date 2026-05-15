package main

import "fmt"

type Puffo struct {
	nome string
	anni int
}

func main() {
	var puffi []Puffo
	puffi = append(puffi, Puffo{"Grande", 542})
	puffi = append(puffi, Puffo{"Quattrocchi", 27})

	fmt.Println("originale:", puffi)

	forrange(puffi)
	fmt.Println("for range:", puffi)

	forternario(puffi)
	fmt.Println("for ternario:", puffi)
}

func forrange(puffi []Puffo) {
	for _, p := range puffi {
		p.anni = 1
	}

	// // equivalente a fare:
	// for i := 0; i < len(puffi); i++ {
	// 	p := puffi[i]
	// 	p.anni = 1

	// 	// puffi[i] = p // QUESTA COSA NON VIENE FATTA
	// }
}

func forternario(puffi []Puffo) {
	for i := 0; i < len(puffi); i++ {
		puffi[i].anni = 1
	}
}
