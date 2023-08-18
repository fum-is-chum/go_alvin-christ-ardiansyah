package main

import "fmt"

/*
  Fungsi pow() dengan time complexity O(log n)
*/
func pow(x, n int64) int64 {
    res := int64(1)
    base := x

    for n > 0 {
        if n%2 == 1 {
            res *= base
        }
        base *= base
        n /= 2
    }

    return res
}

func main() {
	fmt.Println(pow(2, 3)) // 8

    fmt.Println(pow(5, 3)) // 125

    fmt.Println(pow(10, 2)) // 100

    fmt.Println(pow(2, 5)) // 32

    fmt.Println(pow(7, 3)) // 343
}