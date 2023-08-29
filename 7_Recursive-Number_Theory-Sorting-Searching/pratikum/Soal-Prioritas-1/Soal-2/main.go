package main

import (
	"fmt"
	"sort"
)


type pair struct {
	name string
	count int
}
type By func(p1, p2 *pair) bool

func (by By) Sort(items []pair) {

}

func MostAppearItem(items []string) []pair {
	countItemAppearance := make(map[string]int)

	// map items
	for _,item:= range items {
		countItemAppearance[item]++
	}

	// create slices of pair
	var result []pair

	for key,value := range countItemAppearance {
		result = append(result, pair{key, value})
	}

	// sort by appearance count
	sort.Slice(result, func(i,j int) bool {
		return result[i].count < result[j].count
	})

	return result
}

func main() {

	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}) // golang->1 ruby->2 js->4

	for _, list := range pairs {

		fmt.Print(list.name, " -> ", list.count, " ")

	}

	fmt.Println()

	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}) // C->1 D->2 B->3 A->4

	for _, list := range pairs {

		fmt.Print(list.name, " -> ", list.count, " ")

	}

	fmt.Println()

	pairs = MostAppearItem([]string{"football", "basketball", "tenis"}) // football->1 basketball->1 tenis->1

	for _, list := range pairs {

		fmt.Print(list.name, " -> ", list.count, " ")

	}

}
