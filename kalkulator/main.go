package main

import (
	"fmt"
)

func main() {
	fmt.Println("selamat datang di aplikasi kalkulator sederhana")
	fmt.Println("------------------------------------------------")

	var num1, num2 float64
	var operator string

	fmt.Print("Masukan angka pertama: ")
	fmt.Scanln(&num1)

	fmt.Print("Mauskan Operator (+, -, *, /):")
	fmt.Scanln(&operator)

	fmt.Print("Masukan angka kedua: ")
	fmt.Scanln(&num2)

	result := 0.0

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Error: Pemabgian dengan nol tidak valid.")
			return
		}
	default:
		fmt.Println("Error: Operator tidak valid.")
		return

	}
	fmt.Printf("Hasil: %d %s %d = %d\n", num1, operator, num2, result)
}
