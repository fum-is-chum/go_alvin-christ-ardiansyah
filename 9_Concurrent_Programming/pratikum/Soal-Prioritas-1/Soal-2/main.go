package main

import (
	"fmt"
	"sync"
)

func main() {
	resultCh := make(chan int)

	var wg sync.WaitGroup

	// mencetak 50 bilangan pertama dari kelipatan 3
	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go func(x int, ch chan int, wg *sync.WaitGroup){
			defer wg.Done()
			if x % 3 == 0 {
				ch <- x
			}
		}(i, resultCh, &wg)
	}
	
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for result := range resultCh {
		fmt.Println(result)
	}

}
