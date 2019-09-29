package main

import "fmt"

func main() {

}

//WardrobeCombinator wardrobe combinator
func WardrobeCombinator() [][]int {
	sizes := []int{120, 100, 75, 50}
	//a := [][]int{{50, 50, 50, 50, 50}, {100, 100, 50}, {75, 75, 100}}
	sizesAux := [][]int{}
	for _, s := range sizes {
		sizesAux = append(sizesAux, mixer(sizes, s))
	}

	//Idea of possible solution in pseudo-code
	// []int solution1= mi_algoritmo()
	// if(a includes solution1){

	// } else {
	// 	a.add(solution1)
	// }

	return sizesAux
}

func mixer(sizes []int, num int) []int {
	long := 250
	longAux := 0
	sizesAux := []int{}
	for _, i := range sizes {
		longAux = num + i
		if longAux-long >= 50 {
			sizesAux = append(sizesAux, i)
			fmt.Println(sizesAux)
		}
	}

	return sizesAux
}
