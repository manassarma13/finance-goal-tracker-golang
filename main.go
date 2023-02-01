package main

import (
	"fmt"
	"math"
)

func main() {
	monthlyInstallments := 0.0
	goalAmount := 0.0
	goalAge := 0
	currentAge := 0
	currentNetWorth := 0.0
	appreciationValuePercentage := 0.0

	fmt.Print("Enter goal amount: ")
	fmt.Scanf("%f", &goalAmount)

	fmt.Print("Enter goal age: ")
	fmt.Scanf("%d", &goalAge)

	fmt.Print("Enter current age: ")
	fmt.Scanf("%d", &currentAge)

	fmt.Print("Enter current net worth: ")
	fmt.Scanf("%f", &currentNetWorth)

	fmt.Print("Enter appreciation value percentage: ")
	fmt.Scanf("%f", &appreciationValuePercentage)

	months := (goalAge - currentAge) * 12
	amountNeeded := goalAmount - currentNetWorth

	monthlyInstallments = amountNeeded / float64(months)
	monthlyInstallments = monthlyInstallments / math.Pow((1+(appreciationValuePercentage/100)), float64(months/12))

	fmt.Printf("Monthly deposit amount: %.2f", monthlyInstallments)
}
