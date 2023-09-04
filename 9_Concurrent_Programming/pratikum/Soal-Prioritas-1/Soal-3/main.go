package main

import "fmt"

func main() {
	resultCh := make(chan int, 50) // kita ambil 50 bilangan pertama dari kelipatan 3

	// channel hanya aka nmenerima 50 bilangan pertama, tidak peduli seberapa
	// banyak kita mengiterasi i karena channel memiliki ukuran buffer 50
	for i := 1; i <= 100; i++ {
		go func(x int, ch chan int) {
			ch <- 3 * x
		}(i, resultCh)
	}

	fmt.Println(<-resultCh)
}