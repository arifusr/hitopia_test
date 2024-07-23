package main

import "fmt"

var BracketOpenClose map[rune]rune = map[rune]rune{
	'{': '}',
	'(': ')',
	'[': ']',
}

var Brackets map[rune]bool = map[rune]bool{
	'{': true,
	'(': true,
	'[': true,
	'}': true,
	')': true,
	']': true,
}

func BalancedBracket(input string) string {
	listOfBracket := make([]rune, 0)
	for _, v := range input {
		if _, ok := Brackets[v]; ok {
			if len(listOfBracket) > 0 && v == BracketOpenClose[listOfBracket[len(listOfBracket)-1]] {
				listOfBracket = listOfBracket[:len(listOfBracket)-1]
			} else {
				listOfBracket = append(listOfBracket, v)
			}
		}
	}
	if len(listOfBracket) > 0 {
		return "NO"
	} else {
		return "YES"
	}
}

func main() {
	result := BalancedBracket("{ ( ( [ ] ) [ ] ) [ ] }")
	fmt.Print(result)
}
