//На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

package main

import "fmt"

func main() {
	var text, new_text string

	fmt.Scan(&text)

	for i := range text {
		if i%2 != 0 {
			new_text += string(text[i])
		}
	}

	fmt.Printf("%s", new_text)
}
