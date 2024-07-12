package main

import (
	"fmt"
	"log"

	"SF-HW-7.7.1/calc"
)

// userInput reads input from the console for the specified variable and message.
func userInput(variable interface{}, msg string) error {
	fmt.Printf("%s: ", msg)
	_, err := fmt.Scanln(variable)
	return err
}

func main() {
	log.SetFlags(0)

	var (
		firstNum  float64
		secondNum float64
		operator  string
	)

	err := userInput(&firstNum, "Enter first number")
	if err != nil {
		log.Fatal(err)
	}

	err = userInput(&operator, "Enter operator")
	if err != nil {
		log.Fatal(err)
	}

	err = userInput(&secondNum, "Enter second number")
	if err != nil {
		log.Fatal(err)
	}

	c := calc.NewCalculator()
	res, err := c.Calculate(firstNum, secondNum, operator)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v %s %v = %v\n", firstNum, operator, secondNum, res)
}
