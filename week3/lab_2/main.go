package main

import (
	"bufio"
	"fmt"
	"os"

	"crack/utils/crack"
)

func main() {
	target := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"

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

		if crack.CheckSHA1(word) == target {
			fmt.Println("\n=== FOUND ===")
			fmt.Println("Password:", word)
			return
		}
	}

	fmt.Println("\nNot found.")
}
