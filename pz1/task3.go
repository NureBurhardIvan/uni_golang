package main

import (
	"fmt"
)

func convertMassToKg(unit int, mass float64) {
    switch unit {
    case 1:
        fmt.Printf("Маса в кілограмах: %.6f\n", mass/1000000)
    case 2:
        fmt.Printf("Маса в кілограмах: %.6f\n", mass/1000)
    case 3:
        fmt.Printf("Маса в кілограмах: %.6f\n", mass)
    case 4:
        fmt.Printf("Маса в кілограмах: %.6f\n", mass*100)
    case 5:
        fmt.Printf("Маса в кілограмах: %.6f\n", mass*1000)
    default:
        fmt.Println("Невірна одиниця виміру")
    }
}
