package main

import (
	"fmt"
)

func devidable(s *string, d *string) bool {
	sLen := len(*s)
	dLen := len(*d)
	if sLen%dLen != 0 {
		return false
	}

	for i := 0; i < sLen/dLen; i++ {
		for j := 0; j < dLen; j++ {
			if (*d)[j] != (*s)[i*dLen+j] {
				return false
			}
		}

	}
	return true
}

func findAllDevisor(s *string) []*string {
	devisorList := []*string{}

	length := len(*s)
	for i := length; i > 0; i-- {
		if length%i != 0 {
			continue
		}
		dev := string((*s)[0:i])
		if devidable(s, &dev) {
			devisorList = append(devisorList, &dev)
		}
	}
	return devisorList
}

func gcdOfStrings(str1 string, str2 string) string {
	devisorListSortedByLength := findAllDevisor(&str1)
	for _, devisor := range devisorListSortedByLength {
		if devidable(&str2, devisor) == true {
			return *devisor
		}
	}
	return ""

}

func main() {
	str1 := "ABCABC"
	str2 := "ABC"
	fmt.Println(gcdOfStrings(str1, str2))
	str1 = "ABABAB"
	str2 = "ABAB"
	fmt.Println(gcdOfStrings(str1, str2))
	str1 = "LEET"
	str2 = "CODE"
	fmt.Println(gcdOfStrings(str1, str2))
	str1 = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	str2 = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	fmt.Println(gcdOfStrings(str1, str2))
}
