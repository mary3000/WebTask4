package main

import (
	//"fmt"
	"unicode"
	"strings"
)

func RemoveEven(array []int) []int {
	answer := make([]int, 0, len(array))
	for i := 0; i < len(array); i++ {
		if array[i] % 2 == 1 {
			answer = append(answer, array[i])
		}
	}
	return answer
}

func PowerGenerator(x int) (func() int) {
	num := x
	pow := 1
	return func() int {
		pow = pow*num
		return pow
	}
}

func DifferentWordsCount(s string) int {
	dict := make(map[string]bool)
	begin := -1
	flag := false
	for i := 0; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) {
			flag = true
		} else if flag {
			_, ok := dict[strings.ToLower(s[begin+1:i])]
			if !ok {
				dict[strings.ToLower(s[begin+1:i])] = true
			}
			flag = false
			begin = i
		} else {
			begin = i
		}
	}
	return len(dict)
}
/*
func main() {
	a := []int{0, 3, 2, 5}
	b := RemoveEven(a)
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}

	gen := PowerGenerator(2)
	fmt.Println(gen()) // Должно напечатать 2
	fmt.Println(gen()) // Должно напечатать 4
	fmt.Println(gen()) // Должно напечатать 8

	fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))
	// Должно напечатать 2
}
*/