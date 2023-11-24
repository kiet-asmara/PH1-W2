package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Area          float64
	Circumference float64
}

func areaCirc(diameter float64, chOut chan Circle) {

	area := math.Pi * (diameter / 2) * (diameter / 2)
	circ := 2 * math.Pi * (diameter / 2)

	chOut <- Circle{Area: area, Circumference: circ}

}

func main() {
	diameters := []float64{1, 2, 3, 4, 5, 6}
	chOut := make(chan Circle)

	for _, v := range diameters {
		go areaCirc(v, chOut)
	}
	for i := 0; i <= len(diameters); i++ {
		res := <-chOut
		fmt.Println("output:", res)
	}
}
