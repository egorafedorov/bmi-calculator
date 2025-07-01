package main

import (
	"fmt"
	"math"

	"github.com/fatih/color"
)

const BMIPower = 2

func main() {
	fmt.Println("*** BMI Calculator ***")
	for {
		userWeight, userHeight := getUserInput()
		BMI := calculateBMI(userWeight, userHeight)
		outputResult(BMI)
		userRepeatCalculation := userRepeatCalculation()
		if !userRepeatCalculation {
			break
		}
	}
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64
	fmt.Print("Enter your height in centimeters: ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight in kilograms: ")
	fmt.Scan(&userWeight)
	return userWeight, userHeight
}

func calculateBMI(userWeight float64, userHeight float64) float64 {
	BMI := userWeight / math.Pow(userHeight/100, BMIPower)
	return BMI
}

func outputResult(BMI float64) {
	result := fmt.Sprintf("Your body mass index: %.0f", BMI)
	color.Blue(result)
	switch true {
	case BMI < 16:
		color.Red("You are severely underweight")
	case BMI < 18.5:
		color.Yellow("You are underweight")
	case BMI < 25:
		color.Green("You are of normal weight")
	case BMI < 30:
		color.Yellow("You are overweight")
	default:
		color.Red("You have a degree of obesity")
	}
}

func userRepeatCalculation() bool {
	var userChoice string
	fmt.Print("Do you want to repeat the calculation? (Y/N): ")
	fmt.Scan(&userChoice)
	if userChoice == "Y" || userChoice == "y" {
		return true
	} else {
		return false
	}
}
