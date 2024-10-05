package main

import "fmt"

func main() {
	var N int
	var unit int
	var mass float64

	fmt.Println("ЗАВДАННЯ 1 (20 варіант)")
	fmt.Println("Введіть чотиризначне число:")
	fmt.Scan(&N)

	if IsGeometricProgression(N) {
		fmt.Println("Цифри числа є геометричною прогресією")
	} else {
		fmt.Println("Цифри числа не є геометричною прогресією")
	}

	fmt.Println("ЗАВДАННЯ 3 (20 варіант)")
	fmt.Println("Оберіть одиницю виміру маси:")
    fmt.Println("1 – міліграм, 2 – грам, 3 – кілограм, 4 – центнер, 5 – тонна")
    fmt.Scan(&unit)

    fmt.Println("Введіть масу:")
    fmt.Scan(&mass)

	convertMassToKg(unit, mass)

	fmt.Println("ЗАВДАННЯ 5 (20 варіант)")
	  matrix := [][]float64{
        {1, -2, 3, 4, -5},
        {-6, 7, -8, 9, 10},
        {-11, -12, 13, -14, -15},
        {16, -17, 18, -19, -20},
        {-21, 22, 23, -24, 25},
    }
    
    calculateGeometricMean(matrix)
}


//go run main.go task1.go task3.go task5.go - to run the project