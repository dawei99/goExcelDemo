package main

import (
	"strings"
)

// 数字转字母
func num2Abc(num int) string {
	var zf = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	if num > 26 {
		base := num / 26
		remainder := num % 26
		str := strings.Repeat("A", base) + zf[remainder]
		return str
	}
	return zf[num]
}
