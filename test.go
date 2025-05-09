package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("что хотит найти?")
	a := ""
	d := 0.0
	fmt.Scan(&a)
	transaction := []float64{}

	for i := 0; ; i++ {
		fmt.Scan(&d)
		if d == 0 {
			break
		} else {
			transaction = append(transaction, d)
		}

	}
	fmt.Println(transaction)

	if a == "avg" {
		avg := 0.0
		for _, value := range transaction {
			avg += value

		}
		avg1 := 0.0
		for index := range transaction {
			avg1 += float64(index)

		}
		fmt.Println(avg / avg1)

	} else if a == "sum" {
		sum := 0.0
		for _, value := range transaction {
			sum += value
			fmt.Println(sum)

		}

	} else if a == "med" {
		sort.Float64s(transaction)
		fmt.Println(transaction)

		if len(transaction)%2 == 0 { // кол-во чисел четное
			midle2 := 0
			midle3 := 0
			midle2 = len(transaction) / 2
			midle3 = midle2 - 1

			fmt.Println((transaction[midle2] + transaction[midle3]) / 2)

		} else {
			midle1 := 0
			midle1 = len(transaction) / 2

			fmt.Println(transaction[midle1])

		}

	}
}
