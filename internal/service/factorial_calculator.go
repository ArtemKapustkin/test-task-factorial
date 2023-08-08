package service

import (
	"sync"
)

type FactorialCalculator struct {
}

func NewFactorialCalculator() *FactorialCalculator {
	return &FactorialCalculator{}
}

func (c *FactorialCalculator) factorial(num int) uint64 {
	var result uint64 = 1

	for num > 1 {
		result *= uint64(num)
		num -= 1
	}

	return result
}

func (c *FactorialCalculator) Calculate(a, b int) (uint64, uint64) {
	var wg sync.WaitGroup
	wg.Add(2)

	var factorialA, factorialB uint64

	go func() {
		factorialA = c.factorial(a)
		wg.Done()
	}()

	go func() {
		factorialB = c.factorial(b)
		wg.Done()
	}()

	wg.Wait()

	return factorialA, factorialB
}
