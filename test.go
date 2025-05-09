package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("что хотите найти?")
	a := ""
	d := 0.0
	fmt.Scan(&a)
	transaction := []float64{}

	fmt.Println("Введите числа(0 для выхода)")
	for i := 0; ; i++ {

		fmt.Scan(&d)
		fmt.Println("Введите числа(0 для выхода)")
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
		fmt.Println("среднее значение =", avg/avg1)

	} else if a == "sum" {
		sum := 0.0
		for _, value := range transaction {
			sum += value

		}
		fmt.Println("сумма =", sum)

	} else if a == "med" {
		sort.Float64s(transaction)
		fmt.Println(transaction)

		if len(transaction)%2 == 0 { // кол-во чисел четное
			midle2 := 0
			midle3 := 0
			midle2 = len(transaction) / 2
			midle3 = midle2 - 1

			fmt.Println("медиана чисел =", (transaction[midle2]+transaction[midle3])/2)

		} else {
			midle1 := 0
			midle1 = len(transaction) / 2
			fmt.Println("медиана чисел =", transaction[midle1])

		}

	} else {

		error := "Вы ввели что-то не то, попробуйте снова"
		fmt.Println(error)

	}

}
