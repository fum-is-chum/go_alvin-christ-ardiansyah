package main

import "fmt"

func main() {
	var a,b,h float32;
	fmt.Println("Hitung Luas Trapesium")
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai h: ")
	fmt.Scan(&h)

	fmt.Println(fmt.Sprintf("Luas: %v", HitungLuasTrapesium(a,b,h)))
}

func HitungLuasTrapesium(a,b,h float32) float32 {
	return (a+b) * h / 2;
}