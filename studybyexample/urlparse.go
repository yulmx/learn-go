package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {

	// 包含了一个 scheme，认证信息，主机名，端口，路径，查询参数和片段
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, e := url.Parse(s)
	if e != nil {
		panic(e)
	}

	fmt.Println("Scheme\t", u.Scheme)
	fmt.Println("User\t", u.User)
	fmt.Println("Username\t", u.User.Username())
	p, _ := u.User.Password()
	fmt.Println("Password\t", p)

	fmt.Println("Host\t", u.Host)
	h := strings.Split(u.Host, ":")
	fmt.Println(h[0])
	fmt.Println(h[1])

	fmt.Println("Path\t", u.Path)
	fmt.Println("Fragment\t", u.Fragment)
	fmt.Println("RawQuery\t", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

}