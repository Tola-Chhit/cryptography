package main

import (
	"bufio"
	"fmt"
	"os"

	"crack/utils/crack"
)

func main() {
	target := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f0" +
		"26ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"

	file, err := os.Open("wordlists/nord_vpn.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 0

	for scanner.Scan() {
		line++
		word := scanner.Text()

		fmt.Printf("[%d] %s\n", line, word)

		if crack.CheckSHA512(word) == target {
			fmt.Println("\n=== FOUND ===")
			fmt.Println("Password:", word)
			return
		}
	}

	fmt.Println("\nNot found.")
}
