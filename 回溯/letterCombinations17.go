package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	var result []string
	if digits == "" {
		return result
	}
	table := [][]string{{},{},{"a","b","c"}, {"d","e","f"},
		{"g","h","i"}, {"j","k","l"},
		{"m","n","o"}, {"p","q","r","s"},
		{"t","u","v"}, {"w","x","y","z"}}

	//len := len(digits)
	i := 0
	temp := ""
	backTrack2(&result,table,i,digits,temp )
	return result
}

func backTrack2(result *[]string, table [][]string, i int, digits string, temp string) {
	if len(temp) == len(digits) {
		*result=append(*result,temp)
		temp = ""
		return
	}

	nums := int([]byte(digits)[i]-48)
	key := table[nums]

	for j := 0 ; j < len(key) ; j++ {
		temp += key[j]
		backTrack2(result,table,i+1,digits,temp)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	pho := "233"
	fmt.Print(letterCombinations(pho))
}
