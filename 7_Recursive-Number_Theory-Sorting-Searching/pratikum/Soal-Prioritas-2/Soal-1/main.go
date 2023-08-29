package main

import "fmt"

func isCardMatchDeck(card []int, deck []int) bool {
	return card[0] == deck[0] || card[0] == deck[1] || card[1] == deck[0] || card[1] == deck[1]
}

func playingDomino(cards [][]int, deck []int) interface{} {
	var result []int
	for _,card := range cards {
		if isCardMatchDeck(card, deck) {
			result = card
			break
		}
	}

	if result == nil {
		return "Tutup kartu"
	}
	return result
}

func main() {

	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3})) // [3, 4]

	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6})) // [6, 5]

	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1})) // "Tutup kartu"

}
