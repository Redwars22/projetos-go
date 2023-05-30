package main

import "fmt"

func sum(x float32, y float32) float32 {
	return x + y
}

func subtract(x float32, y float32) float32 {
	return x - y
}

func multiply(x float32, y float32) float32 {
	return x * y
}

func divide(x float32, y float32) float32 {
	return x / y
}

func calculate(x float32, y float32, operator string) {
	switch operator {
	case "+":
		sum(x, y)
	case "-":
		subtract(x, y)
	case "/":
		multiply(x, y)
	case "*":
		divide(x, y)
	default:
		fmt.Println("Insira um operador válido!!!")
	}
}

func menu() {
	var num1 float32
	var num2 float32
	var operator string

	fmt.Print("Insira um número: ")
	fmt.Scanln(&num1)
	fmt.Print("Insira um operador matemático (+ - / *): ")
	fmt.Scanln(&operator)
	fmt.Print("Insira um outro número: ")
	fmt.Scanln(&num2)

	calculate(num1, num2, operator)
}

func main() {
	fmt.Println("--------- ANDREW'S AWESOME CALCULATOR ---------")
	menu()
}
