package main

import (
	"fmt"
	"strings"
)

func reverse(word string) string {
	reversed := ""
	for i := len(word) - 1; i >= 0; i-- {
		reversed += string(word[i])
	}
	return reversed
}

func main() {
	fmt.Println("Program Pembalik Kata")

	for {
		fmt.Println("Masukkan kalimat dengan minimal 3 kata:")

		var input string
		fmt.Scanln(&input) // Mengambil input satu baris

		words := strings.Fields(input)

		if len(words) < 3 {
			fmt.Println("Input kurang dari 3 kata. Silakan coba lagi.")
			return
		}

		fmt.Print("Hasil: ")
		var reversedWords []string
		for _, word := range words {
			reversedWords = append(reversedWords, reverse(word))
		}
		fmt.Println(strings.Join(reversedWords, " "))
		break
	}
}
