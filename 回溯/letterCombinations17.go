package main

import "math"

func letterCombinations(digits string) []string {
	result := make([]string, math.Pow(4, float64(len(digits))))
	if digits == "" {
		return result
	}
	var table []string = []string {"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	//len := len(digits)
	i := 0
	temp := ""
	backTrack(&result,table,i,digits,temp )
	return result
}

func backTrack(result *[]string, table []string, i int, digits string, temp string) {
	if len(temp) == len(digits) {
		result[]
		temp = ""
	}

	nums := digits[i]
	key := table[nums]

	for j := 0 ; j < len(key) ; j++ {
		append(temp,key[j])

	}


}

func main() {
	
}
