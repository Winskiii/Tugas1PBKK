package main

import (
	"bufio"
	"fmt"
	"os"
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
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Input minimal 3 kata:")

		scanner.Scan()
		text := scanner.Text()
		parts := strings.Fields(text)

		if len(parts) < 3 {
			fmt.Println("Input < 3 kata, Masukkan lagi ")
			continue
		}

		fmt.Print("Output: ")
		for _, parts := range parts {
			fmt.Print(reverse(parts) + " ")
		}
		fmt.Println()
		break
	}
}
