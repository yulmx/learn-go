package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := "abc123!?$*&()'-=@~"

	senc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(senc)

	sdec, _ := base64.StdEncoding.DecodeString(senc)
	fmt.Println(string(sdec))

	uenc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uenc)
	udec, _ := base64.URLEncoding.DecodeString(uenc)
	fmt.Println(string(udec))


}