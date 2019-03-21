package main

import "fmt"

// Global constants
const (
	PATH = "./main.go"
	PASS = "secretstuff"
)

func main() {
	fmt.Println("Hashing a small file ....")
	smallhash(PATH)

	fmt.Println("\nHashing a big file ....")
	bighash(PATH)

	fmt.Println("\nHMAC Password Store ....")
	passtore(PASS)

	fmt.Println("\nCSPRNG Generator ....")
	randnum()

	fmt.Println("\nAES Crypto ....")
	adves(PASS)
}

func handle(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
