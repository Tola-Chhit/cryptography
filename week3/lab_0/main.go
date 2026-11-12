package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/sha3"
)

func hash(s string, algo string) string {
	switch algo {
	case "MD5":
		h := md5.Sum([]byte(s))
		return hex.EncodeToString(h[:])
	case "SHA1":
		h := sha1.Sum([]byte(s))
		return hex.EncodeToString(h[:])
	case "SHA256":
		h := sha256.Sum256([]byte(s))
		return hex.EncodeToString(h[:])
	case "SHA512":
		h := sha512.Sum512([]byte(s))
		return hex.EncodeToString(h[:])
	case "SHA3":
		h := sha3.Sum256([]byte(s))
		return hex.EncodeToString(h[:])
	}
	return ""
}

func main() {
	var input1, input2 string

	fmt.Println("======== Name + Hashing Program ========")
	fmt.Print("Please input value 1: ")
	fmt.Scanln(&input1)
	fmt.Print("Please input value 2: ")
	fmt.Scanln(&input2)

	algorithms := []string{"MD5", "SHA1", "SHA256", "SHA512", "SHA3"}

	fmt.Println()
	for _, algo := range algorithms {
		outA := hash(input1, algo)
		outB := hash(input2, algo)

		status := "No Match!"
		if outA == outB {
			status = "Match!"
		}

		fmt.Printf("Hash (%s) : %s / %s => %s\n", algo, outA, outB, status)
	}
}
