package main

import "fmt"

var a float64
var b float64
var action string

func main() {
	defer func ()  {
		fmt.Println("")
		fmt.Println("end program")
	}()

	calculator()
}

func calculator() {
	fmt.Print("Введите первое число... ")
	fmt.Scan(&a)
	fmt.Print("Введите знак действия... ")
	fmt.Scan(&action)
	fmt.Print("Введите второе число... ")
	fmt.Scan(&b)
	fmt.Println("")

	switch action {
	case "+":
		amount(a,b)
	case "-":
		subtraction(a,b)
	case "*":
		multiply(a,b)
	case "/":
		divide(a,b)
	}
}

func amount(a float64, b float64) {
	sum := a+b
	fmt.Println("Ответ ")
	fmt.Printf("%.2f\n", sum)
}

func subtraction(a float64, b float64) {
	sub := a-b
	fmt.Println("Ответ ")
	fmt.Printf("%.2f\n", sub)
}

func multiply(a float64, b float64) {
	mul := a*b
	fmt.Println("Ответ ")
	fmt.Printf("%.2f\n", mul)
}

func divide(a float64, b float64) {
	div := a/b
	fmt.Println("Ответ ")
	fmt.Printf("%.2f\n", div)
}

// норм калькулятор