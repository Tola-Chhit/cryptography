package main

import (
	"bufio"
	"fmt"
	"os"

	"crack/utils/crack"
)

func main() {
	target := "6a85dfd77d9cb35770c9dc6728d73d3f"

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

		if crack.CheckMD5(word) == target {
			fmt.Println("\n=== FOUND ===")
			fmt.Println("Password:", word)
			return
		}
	}

	fmt.Println("\nNot found.")
}
