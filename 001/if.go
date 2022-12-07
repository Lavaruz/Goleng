package main

import "fmt"

func main() {
	var saya = [...]string{
		"Assami",
	}
	if saya[0] == "Assami" {
		fmt.Println("Btul")
	} else {
		fmt.Println("salah")
	}
}
