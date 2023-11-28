/*
	На вход подается строка. Нужно определить, является ли она палиндромом.
	Если строка является палиндромом - вывести Палиндром иначе - вывести Нет.
	Все входные строчки в нижнем регистре.
*/

package main

import "fmt"

func main() {
	var text string
	fmt.Scan(&text)
	var reversedText string
	for _, v := range text {
		reversedText = string(v) + reversedText
	}

	if text == reversedText {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
