package main

import "fmt"

type Car struct {
	carType string
	fuelIn float64	
}

func (c Car) calculateEstDistance() float64 {
	const distancePerFuel = 1.5
	return c.fuelIn * distancePerFuel
}

func main() {
	myCar := Car{}

	fmt.Print("fuel in: ")
	fmt.Scanln(&myCar.fuelIn)

	fmt.Print("car type: ")
	fmt.Scanln(&myCar.carType)


	fmt.Println(fmt.Sprintf("car type: %s, est. distance: %v", myCar.carType, myCar.calculateEstDistance()))
}