//Даются две строки X и S.
//Нужно найти и вывести первое вхождение подстроки S в строке X.
//Если подстроки S нет в строке X - вывести -1

package main

import (
	"fmt"
	"strings"
)

func main() {
	var X, S string
	fmt.Scan(&X)
	fmt.Scan(&S)

	fmt.Println(strings.Index(X, S))

}
