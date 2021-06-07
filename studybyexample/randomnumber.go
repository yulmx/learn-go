package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Float64())

	fmt.Println(rand.Float64() * 5 + 5)
	fmt.Println(rand.Float64() * 5 + 5)


	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Print(r1.Intn(100), "\t", r1.Intn(100), "\t", r1.Intn(100), "\n")

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), "\t", r2.Intn(100), "\t", r2.Intn(100), "\n")
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), "\t", r3.Intn(100), "\t", r3.Intn(100), "\n")

	fmt.Print(r1.Intn(100), "\t", r1.Intn(100), "\t", r1.Intn(100), "\n")

}