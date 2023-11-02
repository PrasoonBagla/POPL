package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"
)

func fib(n int) *big.Int {
	calcFactor := [][]*big.Int{
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(1), big.NewInt(0)},
	}
	mat := [][]*big.Int{
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(1), big.NewInt(0)},
	}
	for i := 1; i < n; i++ {
		matNew := [][]*big.Int{
			{big.NewInt(0), big.NewInt(0)},
			{big.NewInt(0), big.NewInt(0)},
		}
		matNew[0][0].Add(
			big.NewInt(1).Mul(mat[0][0], calcFactor[0][0]),
			big.NewInt(1).Mul(mat[0][1], calcFactor[1][0]),
		)
		matNew[0][1].Add(
			big.NewInt(1).Mul(mat[0][0], calcFactor[0][1]),
			big.NewInt(1).Mul(mat[0][1], calcFactor[1][1]),
		)
		matNew[1][0].Add(
			big.NewInt(1).Mul(mat[1][0], calcFactor[0][0]),
			big.NewInt(1).Mul(mat[1][1], calcFactor[1][0]),
		)
		matNew[1][1].Add(
			big.NewInt(1).Mul(mat[1][0], calcFactor[0][1]),
			big.NewInt(1).Mul(mat[1][1], calcFactor[1][1]),
		)
		mat = matNew
	}
	return mat[0][1]
}

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	startTime := time.Now() // Capture the start time

	res := fib(num)

	duration := time.Since(startTime) // Compute the duration

	fmt.Println(res.String())
	fmt.Printf("Execution time: %s\n", duration)
}