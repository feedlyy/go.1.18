package go_1_18

import (
	"fmt"
	"testing"
)

// CountTotal is used to count the total for a slice of data type int64, int or float64
func CountTotal[V int64 | int | float64](points []V) V {
	var total V

	for _, value := range points {
		total += value
	}
	return total
}

func TestGenerics(t *testing.T) {
	intvalues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	int64values := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	float64values := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}

	fmt.Println("Total for generic ints:", CountTotal[int](intvalues))
	fmt.Println("Total for generic int64:", CountTotal[int64](int64values))
	fmt.Println("Total for generic floats:", CountTotal[float64](float64values))
}
