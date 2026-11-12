package main

import (
	"fmt"
	"regexp"
)

func main() {
	flag := "cryptoCTF{meowmeow}"

	re := regexp.MustCompile(`^cryptoCTF\{(?:\x6d\x65\x6f\x77){2}\}$`)

	if re.MatchString(flag) {
		fmt.Println("Flag is valid:", flag)
	} else {
		fmt.Println("Flag is invalid")
	}
}
