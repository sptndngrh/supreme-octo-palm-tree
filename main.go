package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sptndngrh/supreme-octo-palm-tree/calculations"
)

func readFloatInput(reader *bufio.Reader, prompt string) (float64, error) {
	fmt.Print(prompt)
	inputStr, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	inputStr = strings.TrimSpace(inputStr)
	num, err := strconv.ParseFloat(inputStr, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("========================")
		fmt.Println("Pilih jenis perhitungan:")
		fmt.Println("1. Penjumlahan")
		fmt.Println("2. Pengurangan")
		fmt.Println("3. Perkalian")
		fmt.Println("4. Pembagian")
		fmt.Println("5. Kuadrat")
		fmt.Println("6. Keluar")
		fmt.Print("Pilihan: ")

		choiceStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Terjadi kesalahan saat membaca pilihan")
			return
		}
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Pilihan tidak valid")
			continue
		}

		if choice == 6 {
			fmt.Println("Program berakhir. Sampai jumpa!")
			break
		}

		var result float64

		switch choice {
		case 1, 2, 3, 4:
			num1, err := readFloatInput(reader, "Masukkan angka pertama: ")
			if err != nil {
				fmt.Println("Terjadi kesalahan saat membaca angka pertama")
				continue
			}
			num2, err := readFloatInput(reader, "Masukkan angka kedua: ")
			if err != nil {
				fmt.Println("Terjadi kesalahan saat membaca angka kedua")
				continue
			}

			switch choice {
			case 1:
				result = calculations.Add(num1, num2)
			case 2:
				result = calculations.Subtract(num1, num2)
			case 3:
				result = calculations.Multiply(num1, num2)
			case 4:
				result = calculations.Divide(num1, num2)
			}
		case 5:
			num, err := readFloatInput(reader, "Masukkan angka: ")
			if err != nil {
				fmt.Println("Terjadi kesalahan saat membaca angka")
				continue
			}
			result = calculations.Square(num)
		default:
			fmt.Println("Pilihan tidak valid")
			continue
		}

		fmt.Printf("Hasil: %.2f\n", result)
	}
}
