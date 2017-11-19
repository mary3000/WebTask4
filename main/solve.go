package main

import (
	//"fmt"
	"unicode"
	"strings"
	"fmt"
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
	if flag {
		_, ok := dict[strings.ToLower(s[begin+1:len(s)])]
		if !ok {
			dict[strings.ToLower(s[begin+1:len(s)])] = true
		}
	}
	return len(dict)
}

func main() {
	fmt.Println(DifferentWordsCount("Hello, w world!HELLO  wOrlD..w.12wd"))
	// Должно напечатать 2
}
