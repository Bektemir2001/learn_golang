/*
	байтовый срез можно изменять (в отличие от строки);

	к байтовому срезу применимо все, что применимо к массивам и срезам других типов (взятие среза, итерация);

	к отдельным байтам применимы операции, применимые к числам.
*/

package main

import "fmt"

func main() {
	byteV := []byte("Byte string")

	fmt.Println(byteV)
	fmt.Printf("%s \n", byteV)
	for i := range byteV {
		if byteV[i]%2 == 0 {
			continue
		}
		byteV[i] += 1
	}

	fmt.Println(byteV)
	fmt.Printf("%s \n", byteV)
}
