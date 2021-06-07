package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println("regexp.MatchString\t", match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println("MatchString\t", r.MatchString("peach"))
	fmt.Println("FindString\t", r.FindString("peach punch"))
	fmt.Println("FindStringIndex\t", r.FindStringIndex("peach punch"))
	fmt.Println("FindStringSubmatch\t", r.FindStringSubmatch("peach punch"))
	fmt.Println("FindStringSubmatchIndex\t", r.FindStringSubmatchIndex("peach punch"))

	fmt.Println("FindAllString\t", r.FindAllString("peach punch pinch", -1))
	fmt.Println("FindAllString\t", r.FindAllString("peach punch pinch", 2))
	fmt.Println("FindAllStringSubmatchIndex\t", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Println("Match([]byte)\t", r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("MustCompile\t", r)
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println("ReplaceAllFunc", string(out))

}