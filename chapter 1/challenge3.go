package main

import "fmt"

func main() {
	var word string = "selamat malam"
	m := make(map[string]int)

	for _, char := range word {
		str := string(char)
		fmt.Println(str)
		m[str] = m[str] + 1
	}
	fmt.Println(m)
}
