package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Элементы числовой последовательности Фибоначчи:")

	f1 := big.NewInt(0)
	fmt.Print(f1, " ")

	f2 := big.NewInt(1)

	for i := 1; i <= 99; i++ {
		f1.Add(f1, f2)
		f1, f2 = f2, f1
		fmt.Print(f1, " ")
	}
}
