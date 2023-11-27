package main

import (
	"fmt"
	"math"
)

const (
	RECTANGLE = "rectangle"
	CIRCLE    = "circle"
	TRIANGLE  = "triangle"
)

type Shape struct {
	ShapeType string
	Length    int
	Area      float32
}

func main() {
	chRec := make(chan Shape)
	chCir := make(chan Shape)
	chTri := make(chan Shape)

	input := []Shape{
		{ShapeType: RECTANGLE, Length: 5},
		{ShapeType: CIRCLE, Length: 3},
		{ShapeType: TRIANGLE, Length: 5},
		{ShapeType: RECTANGLE, Length: 15},
		{ShapeType: CIRCLE, Length: 5},
	}

	// assign shapes to channels
	for _, s := range input {
		go func(s Shape) {
			switch s.ShapeType {
			case RECTANGLE:
				chRec <- s
			case CIRCLE:
				chCir <- s
			case TRIANGLE:
				chTri <- s
			}
		}(s)
	}

	// perform operation once channel receives value
	for i := 0; i < len(input); i++ {
		select {
		case s := <-chRec:
			s.Area = float32(s.Length) * float32(s.Length)
			fmt.Println(s)
		case s := <-chCir:
			s.Area = math.Pi * float32(s.Length) * float32(s.Length)
			fmt.Println(s)
		case s := <-chTri:
			s.Area = 0.5 * float32(s.Length) * float32(s.Length)
			fmt.Println(s)
		}
	}
}
