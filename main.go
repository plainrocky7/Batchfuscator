package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	obfuscate()

}

func genrandomstring(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func obfuscate() {

	file, err := os.OpenFile(os.Args[1], os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Print("Error opening file!")
	}

	fmt.Printf("Obfuscating %v \n", file)
	for i := 0; 402 > i; i++ {

		_, err := file.WriteString("set " + genrandomstring(403) + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

	}
	fmt.Printf("finished obfuscating %v", file)
	file.Close()
}
