package main

import (
	"fmt"
)

func main() {
	//var nums = 746627324245

	fmt.Println(isIsomorphic("egg", "add"))
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS := make(map[byte]byte)
	mapT := make(map[byte]byte)

	for i := range s {
		valS, okS := mapS[s[i]]
		valT, okT := mapT[t[i]]

		if !okS && !okT {
			mapS[s[i]] = t[i]
			mapT[t[i]] = s[i]
		} else if okS {
			if valS != t[i] {
				return false
			}
		} else if okT {
			if valT != s[i] {
				return false
			}
		}
	}

	return true
}
