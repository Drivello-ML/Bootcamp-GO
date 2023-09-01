package main

import (
	"github.com/Drivello-ML/Bootcamp-GO/gobases/functions"
)

func main() {
	var choice int
	var numA, numB float64

	functions.ShowMenu(&choice)
	functions.Ask4Numbers(&numA, &numB)
	functions.DoMath(choice, numA, numB)
}
