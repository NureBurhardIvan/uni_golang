package main

import (
	"fmt"
	"math"
)

func calculateGeometricMean(matrix [][]float64) {
    var negativeModules []float64

    for i := 0; i < len(matrix); i++ {
        if matrix[i][4] < 0 {
            negativeModules = append(negativeModules, math.Abs(matrix[i][4]))
        }
    }

    if len(negativeModules) == 0 {
        fmt.Println("Немає від'ємних елементів у п'ятому стовпці.")
        return
    }

    product := 1.0
    for _, val := range negativeModules {
        product *= val
    }

    geometricMean := math.Pow(product, 1.0/float64(len(negativeModules)))
    fmt.Printf("Середнє геометричне модулів від'ємних елементів п'ятого стовпця: %.6f\n", geometricMean)
}
