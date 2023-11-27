package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("test sentence", "test"))
	fmt.Println(strings.Count("Hi Guys", "e"))
	fmt.Println(strings.HasPrefix("My book", "m"))
	fmt.Println(strings.HasSuffix("Finishing course", "e"))
	fmt.Println(strings.Index("Some text", "te"))
	fmt.Println(strings.Join([]string{"start", "end"}, "-"))

	// Повторяет строку n раз подряд
	fmt.Println(strings.Repeat("a", 5))

	// Функция Replace заменяет любое вхождение old в вашей строке на new
	// Если значение n равно -1, то будут заменены все вхождения.
	// Общий вид: func Replace(s, old, new string, n int) string
	// Пример:
	fmt.Println(strings.Replace("blanotblanot", "not", "***", -1))

	// Разбивает строку согласно разделителю
	fmt.Println(strings.Split("a-b-c-d-e", "-"))

	// Возвращает строку c нижним регистром
	fmt.Println(strings.ToLower("TEST"))

	// Возвращает строку c верхним регистром
	fmt.Println(strings.ToUpper("test"))
	// результат: "TEST"

	// Возвращает строку с вырезанным набором
	fmt.Println(strings.Trim("test", "te"))

}
