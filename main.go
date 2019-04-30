package main

import (
	"fmt"
	"math/rand"
)

func main() {

	X := makeX()
	Y := makeY(X)

	fmt.Println(Y)
	fmt.Println(len(Y))
}

func makeX() []float64 {
	X := []float64{}

	n := 100
	for i := 0; i < n; i++ {
		X = append(X, rand.Float64())
	}
	return X
}

func makeY(X []float64) []float64 {
	Y := []float64{}

	for _, x := range X {
		std := 0.01
		Y = append(Y, x+rand.NormFloat64()*std)

	}
	return Y

}
