package example

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := float64(1)

	prevz := float64(-1)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %d -- %.15f\n", i, z)
		if (math.Round(prevz/0.00005) * 0.00005) == (math.Round(z/0.00005) * 0.00005) {
			return z
		}
		prevz = z
	}
	return z
}

func sqrtRun() {
	// Exercise details can be found at https://tour.golang.org/flowcontrol/8
	fmt.Println(sqrt(2))
}

func init() {
	category := GetCategories().AddCategory("flow_control")

	category.AddExample("square_root",
		CategoryExample{
			Description: "Finding a square root?",
			Function:    sqrtRun,
		})
}
