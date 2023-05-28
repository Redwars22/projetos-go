package main

import "fmt"

func getDayOfTheWeek(_year int, _month int, _day int) int {
	year := _year
	month := _month
	day := _day

	if month < 3 {
		year = year - 1
		year = year + 12
	}

	a := year / 100
	b := year % 100

	dayOfTheWeek := (day + ((13 * (month + 1)) / 5) + b + (b / 4) + (a / 4) - (2 * a)) % 7

	return dayOfTheWeek
}

func main() {
	var year int
	var day int
	var month int

	fmt.Println("Descubra o dia da semana a partir de uma data")
	fmt.Print("Dia: ")
	fmt.Scanln(&day)
	fmt.Print("Mês: ")
	fmt.Scanln(&month)
	fmt.Print("Ano: ")
	fmt.Scanln(&year)

	dayOfTheWeek := getDayOfTheWeek(year, month, day)

	if dayOfTheWeek == 0 {
		fmt.Println("Sábado")
	} else if dayOfTheWeek == 1 {
		fmt.Println("Domingo")
	} else if dayOfTheWeek == 2 {
		fmt.Println("Segunda-Feira")
	} else if dayOfTheWeek == 3 {
		fmt.Println("Terça-Feira")
	} else if dayOfTheWeek == 4 {
		fmt.Println("Quarta-Feira")
	} else if dayOfTheWeek == 5 {
		fmt.Println("Quinta-Feira")
	} else if dayOfTheWeek == 6 {
		fmt.Println("Sexta-Feira")
	}
}
