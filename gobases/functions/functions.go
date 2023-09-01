package functions

import "fmt"

func ShowMenu(choice *int) {
	const menu string = `
1. Add
2. Substract
3. Multiply
4. Divide
`
	fmt.Print(menu)
	fmt.Println("What do you want to do?")
	fmt.Scanln(choice)
}

func Ask4Numbers(numA *float64, numB *float64) {
	fmt.Println(`Insert the first number`)
	fmt.Scanln(numA)
	fmt.Println(`Insert the second number`)
	fmt.Scanln(numB)
}

func DoMath(operation int, numA, numB float64) {
	switch operation {
	case 1:
		fmt.Printf("Your answer is %f", numA+numB)
	case 2:
		fmt.Printf("Your answer is %f", numA-numB)
	case 3:
		fmt.Printf("Your answer is %f", numA/numB)
	case 4:
		fmt.Printf("Your answer is %f", numA*numB)
	default:
		fmt.Println("I guess i was not clear enough. You choosed a wrong option.")
	}
}
