package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const letters = "abcdefghijklmnopqrstuvwxyz0123456789" // se volete potete aggiungere anche simboli come (?,.:!โฌ]ยง) eccetera...

func GenRandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Created By ๐ฃ๐๐ ๐๐ก๐ค๐๐")
	fmt.Print("Inserisci una password: ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-2]

	guess := ""
	for guess != text {
		guess = GenRandString(len(text))
		fmt.Println(guess)
	}

	fmt.Println("la password รจ:", guess)
}
