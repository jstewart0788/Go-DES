package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var cypherText cypherText
	var decodedText string
	reader := bufio.NewReader(os.Stdin)

	// Get input from user
	fmt.Print("Enter your message: ")
	message, _ := reader.ReadString('\n')
	message = removeCarriageWindows(message)
	fmt.Print("Enter your key(length 8):")
	key, _ := reader.ReadString('\n')
	key = removeCarriageWindows(key)
	if len(key) != 8 {
		fmt.Println("Key is incorrect length. Please enter an eight chracter key")
		os.Exit(1)
	}
	// message := "Okay this is a grand opportunity to test this out"
	// key := "ABCDEFGH"

	chunks := stringToEightByteArray(message)
	for _, chunk := range chunks {
		cypherText = append(cypherText, desEnc(chunk, key))
	}

	cypherText.print()

	for _, chunk := range cypherText {
		decodedText = decodedText + desDec(chunk, key)
	}
	fmt.Println("Decoded Text:", decodedText)

}
