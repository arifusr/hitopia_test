package main

import "fmt"

func GetWeightOfChar() func(rune) int {
	mapOfWeight := make(map[rune]int)
	for i, j := 'a', 1; i <= 'z'; i, j = i+1, j+1 {
		mapOfWeight[rune(i)] = j
	}
	return func(s rune) int {
		return mapOfWeight[s]
	}
}

func WeightedString(input1 string, input2 []int) []string {
	getWeightOfChar := GetWeightOfChar()
	stringAndSubstringWeight := make([]int, 0)
	var lastRune rune
	var lastWeight int
	for _, v := range input1 {
		weight := getWeightOfChar(v)
		if v == lastRune {
			lastWeight = lastWeight + weight
		} else {
			lastWeight = weight
		}
		lastRune = v
		stringAndSubstringWeight = append(stringAndSubstringWeight, lastWeight)
	}
	result := make([]string, 0)
	for _, targetWeight := range input2 {
		var exist bool
	inner:
		for _, weight := range stringAndSubstringWeight {
			if weight == targetWeight {
				exist = true
				break inner
			}
		}
		if exist {
			result = append(result, "YES")
		} else {
			result = append(result, "NO")
		}
	}
	return result
}

func main() {
	result := WeightedString("abbcccd", []int{1, 3, 9, 8})
	fmt.Print(result)
}
