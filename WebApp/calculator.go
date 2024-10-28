package main

import (
	"fmt"
)

type Calculator struct {
}

func (c *Calculator) Add(a, b float64) float64 {
	return a + b
}

func (c *Calculator) Subtract(a, b float64) float64 {
	return a - b
}

func (c *Calculator) Multiply(a, b float64) float64 {
	return a * b
}

func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	calc := &Calculator{}

	sum := calc.Add(5, 3)
	fmt.Printf("Sum: %.2f\n", sum)

	difference := calc.Subtract(10, 4)
	fmt.Printf("Difference: %.2f\n", difference)

	product := calc.Multiply(2, 7)
	fmt.Printf("Product: %.2f\n", product)

	quotient, err := calc.Divide(15, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Quotient: %.2f\n", quotient)
	}
}
