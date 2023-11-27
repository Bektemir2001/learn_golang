package main

import "fmt"

func main() {
	runeV := []rune("this is a rune text")
	fmt.Println(runeV)
	for i := range runeV {
		if runeV[i] == 'i' {
			runeV[i] = 'a'
		}
	}

	fmt.Println(runeV)
}
