package main

import "fmt"

func main() {
	var word string = "САШАРВО"
	for i := 0; i < 5; i++ {
		fmt.Printf("Nilai i = %d\n", i)
	}
	for j := 0; j <= 10; j++ {
		if j == 5 {
			for pos, char := range word {
				fmt.Printf("Character %#U starts at byte position %d\n", char, pos)
			}
		} else {
			fmt.Printf("Nilai j = %d\n", j)
		}
	}
}
