package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sptndngrh/supreme-octo-palm-tree/calculations"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Pilih jenis perhitungan:")
	fmt.Println("1. Penjumlahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Perkalian")
	fmt.Println("4. Pembagian")
	fmt.Println("5. Kuadrat")
	fmt.Print("Pilihan: ")

	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)
	choice, err := strconv.Atoi(choiceStr)
	if err != nil {
		fmt.Println("Pilihan tidak valid")
		return
	}

	var result float64

	switch choice {
	case 1:
		fmt.Print("Masukkan angka pertama: ")
		num1Str, _ := reader.ReadString('\n')
		num1, _ := strconv.ParseFloat(strings.TrimSpace(num1Str), 64)

		fmt.Print("Masukkan angka kedua: ")
		num2Str, _ := reader.ReadString('\n')
		num2, _ := strconv.ParseFloat(strings.TrimSpace(num2Str), 64)

		result = calculations.Add(num1, num2)
	case 2:
		fmt.Print("Masukkan angka pertama: ")
		num1Str, _ := reader.ReadString('\n')
		num1, _ := strconv.ParseFloat(strings.TrimSpace(num1Str), 64)

		fmt.Print("Masukkan angka kedua: ")
		num2Str, _ := reader.ReadString('\n')
		num2, _ := strconv.ParseFloat(strings.TrimSpace(num2Str), 64)

		result = calculations.Subtract(num1, num2)
	case 3:
		fmt.Print("Masukkan angka pertama: ")
		num1Str, _ := reader.ReadString('\n')
		num1, _ := strconv.ParseFloat(strings.TrimSpace(num1Str), 64)

		fmt.Print("Masukkan angka kedua: ")
		num2Str, _ := reader.ReadString('\n')
		num2, _ := strconv.ParseFloat(strings.TrimSpace(num2Str), 64)

		result = calculations.Multiply(num1, num2)
	case 4:
		fmt.Print("Masukkan angka pertama: ")
		num1Str, _ := reader.ReadString('\n')
		num1, _ := strconv.ParseFloat(strings.TrimSpace(num1Str), 64)

		fmt.Print("Masukkan angka kedua: ")
		num2Str, _ := reader.ReadString('\n')
		num2, _ := strconv.ParseFloat(strings.TrimSpace(num2Str), 64)

		result = calculations.Divide(num1, num2)
	case 5:
		fmt.Print("Masukkan angka: ")
		numStr, _ := reader.ReadString('\n')
		num, _ := strconv.ParseFloat(strings.TrimSpace(numStr), 64)

		result = calculations.Square(num)
	default:
		fmt.Println("Pilihan tidak valid")
		return
	}

	fmt.Printf("Hasil: %.2f\n", result)
}

