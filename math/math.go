package math

import (
	"fmt"
	"math"
)

func Math() {
	var name = -123.123
	fmt.Println(math.Abs(name))
}

/**
精度損失測試
*/
func Float() {
	x, y := 0.0, 0.0
	for i := 0; i < 10; i++ {
		x += 0.1
		if i%2 == 0 {
			y += 0.2
		} else {
			fmt.Printf("%-5t %-5t %-5t %-5t", x == y, equalFloat(x, y, -1),
				equalFloat(x, y, 0.000000000001), equalFloat(x, y, 6))
			fmt.Println(x, y)
		}
	}
}

func equalFloat(x, y, limit float64) bool {
	if limit <= 0.0 {
		limit = math.SmallestNonzeroFloat64
	}
	return math.Abs(x-y) <= (limit * math.Min(math.Abs(x), math.Abs(y)))
}
