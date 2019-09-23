package main

type combinator struct {
	in []int
}

func main() {

}

//WardrobeCombinator wardrobe combinator
func WardrobeCombinator() combinator {

	c := combinator{
		in: []int{50, 50, 50, 50, 50},
	}
	return c
}
