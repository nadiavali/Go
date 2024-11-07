package main

import (
	"fmt"
	"strings"
)
func toAlternatingCase(str string) string{
	res := ""
	for _, chr:= range str {
		if string(chr) == strings.ToUpper(string(chr)){
			res += strings.ToLower(string(chr))
		}else {
			res += strings.ToUpper(string(chr))
		}
	}
	return res
}

func main() {
	s := "HeLLo 1WorlD!"
	fmt.Println(toAlternatingCase(s))
}