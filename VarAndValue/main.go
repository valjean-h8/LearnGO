package main

import (
	"fmt"
	"math"
)

func main() {
	// static type this variant better
	// var userWeight float32 = 80.8
	// var userHeight int = 175

	// dynamic type
	// var userWeight = 80.8
	// var userHeight = 1.75

	// implicit definition
	// userWeight := 80.8

	// explict definition
	// var userHeight float64
	// userHeight = 1.75

	//multi defenition
	var userHeight, userWeight float64 = 1.8, 80

	// userHeight, userWeight := 1.8, 80

	// constant
	const IMTPower = 2

	var IMT = userWeight / math.Pow(float64(userHeight), IMTPower)

	fmt.Println(IMT)
}
