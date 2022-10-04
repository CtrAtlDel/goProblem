package main

import (
	"fmt"
)

func merge(s1 string, s2 string) string {
	var s1_size = len(s1)
	var s2_size = len(s2)
	var size = s2_size
	if s1_size-s2_size > 0 {
		size = s1_size
	} else {
		size = s2_size
	}
	var i_i = 0
	var j_j = 0
	var str []byte
	var s1_byte = []byte(s1)
	var s2_byte = []byte(s2)

	for i := 0; i < size; i++ {
		if s1_byte[i_i] >= s2_byte[j_j] {
			str = append(str, s1_byte[j_j])
			j_j++
		} else {
			str = append(str, s1_byte[i_i])
			i_i++
		}

	}
	return string(str)
}

func main() {
	var s1 string
	var s2 string
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)
	fmt.Println(merge(s1, s2))
}
