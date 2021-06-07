package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 读文件
	data, err := ioutil.ReadFile("tmp.txt")
	check(err)
	fmt.Println(string(data))

	f, err := os.Open("tmp.txt")
	check(err)
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s|\n", n1, string(b1))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s|\n", n2, o2, string(b2))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s|\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s|\n", string(b4))

	f.Close()

	// 写文件
	data1 := []byte("hello\ngo\n")
	err = ioutil.WriteFile("tmp2.txt", data1, 0644)
	check(err)

	f, err = os.Create("tmp3.txt")
	check(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err = f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err = f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)
	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffer\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}