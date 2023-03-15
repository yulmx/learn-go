// @program: learn-go
// @author: yulm
// @created: 2023-03-14 18:16
// @note:
package main

import "fmt"

func main() {
	var s = "中国china"
	fmt.Println(len(s), s, s[3])

	// 按字节遍历
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println("___________")

	// 按字符遍历 type rune int32
	for codepoint, runeValue := range s {
		fmt.Printf("%d: %x\n", codepoint, int32(runeValue))
	}
	fmt.Println("___________")
}
