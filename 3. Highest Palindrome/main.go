package main

import (
	"fmt"
	"strconv"
)

func IsPalindrome(s string) bool {
	// split rune
	if len(s) == 1 {
		return true
	}
	L := s[:1]
	R := s[len(s)-1:]
	if L != R {
		return false
	}
	if len(s) > 2 {
		nested := IsPalindrome(s[1 : len(s)-1])
		if nested == false {
			return false
		}
	}

	return true
}

func ToPalindrome(s string, k int) (palindrome [][]string, remainingK int) {
	// split string
	var (
		L, R   string
		nested [][]string
	)
	if len(s) == 1 {
		palindrome = append(palindrome, []string{s})
		return palindrome, k
	} else {
		L = s[:1]
		R = s[len(s)-1:]
		palindrome = append(palindrome, []string{L, R})

	}

	if len(s) > 2 {
		nested, k = ToPalindrome(s[1:len(s)-1], k)
		palindrome = append(palindrome, nested...)
		if k < 0 {
			return palindrome, -1
		}
	}

	var (
		Lint, Rint int
		Lerr, Rerr error
	)
	Lint, Lerr = strconv.Atoi(L)
	Rint, Rerr = strconv.Atoi(R)

	// can be 2 step
	if Lerr != nil && Rerr != nil {
		palindrome[0][0] = "9"
		palindrome[0][1] = "9"
		k -= 2
	} else if Lerr != nil {
		palindrome[0][0] = R
		k -= 1
	} else if Rerr != nil {
		palindrome[0][1] = L
		k -= 1
	} else if Lint > Rint {
		palindrome[0][1] = L
		k -= 1
	} else if Lint < Rint {
		palindrome[0][0] = R
		k -= 1
	}
	return palindrome, k
}

func MaximizedPalindrome(p [][]string, k int) (remainingK int) {
	if k < 2 {
		return k
	}
	if p[0][0] != "9" && p[0][1] != "9" {
		p[0][0] = "9"
		p[0][1] = "9"
		k -= 2
	}
	if k >= 2 {
		k = MaximizedPalindrome(p[1:], k)

	}
	return k
}

func PrintPalindrome(p [][]string) {
	if len(p[0]) == 1 {
		fmt.Print(p[0][0])
		return
	}
	L := p[0][0]
	R := p[0][1]
	fmt.Print(L)
	if len(p) > 1 {
		PrintPalindrome(p[1:])
	}
	fmt.Print(R)
}

func MaximumMiddlePalindrome(p [][]string) {
	if len(p[0]) == 1 {
		p[0][0] = "9"
		return
	}
	if len(p) > 1 {
		MaximumMiddlePalindrome(p[1:])
	}
}

func HighestPalindrome(s string, k int) (higestPalindrome string) {
	p, k := ToPalindrome(s, k)
	if k < 0 {
		return "-1"
	}
	if k >= 2 {
		k = MaximizedPalindrome(p, k)
	}
	if k >= 1 {
		MaximumMiddlePalindrome(p)
	}
	// transform to string
	PrintPalindrome(p)
	// fmt.Println(p, k)
	return ""
}

func main() {
	s := "39443"
	k := 2
	HighestPalindrome(s, k)
}
