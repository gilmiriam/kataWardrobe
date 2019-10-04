package main

func main() {

}

//WardrobeCombinator wardrobe combinator
func WardrobeCombinator() [][]int {
	sizes := []int{100, 75, 50}
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
	// long := 250
	// longAux := 0
	// sizesAux := []int{}
	// total := 0
	// for _, i := range sizes {
	// 	longAux = num + i
	// 	if long-longAux >= 50 {
	// 		total += i
	// 		sizesAux = append(sizesAux, i)
	// 		fmt.Println(sizesAux)
	// 	}
	// }

	// total += num
	// if total == long {
	// 	return sizesAux
	// }

	longTotal := 250
	long := 250
	longAux := 0
	sizesAux := []int{}
	total := 0
	pos := 0
	for total > longTotal || pos < len(sizes) || long < 0 {
		longAux = num + sizes[pos]
		res := longTotal - longAux
		if contains(sizes, res) {
			long = long - sizes[pos]
			total += sizes[pos]
			sizesAux = append(sizesAux, sizes[pos])
		} else {
			pos++
		}
	}
	if total == 250 {
		return sizesAux
	}
	return []int{}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
