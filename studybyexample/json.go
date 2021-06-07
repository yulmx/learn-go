package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type RSP1 struct {
	Page int
	Fruits []string
}

type RSP2 struct {
	Page int	`json:"page"`		// 使用tag来标识json的key，即key有Page变成了page
	Fruits []string	`json:"fruits"`
}

func main() {

	b, _ := json.Marshal(true)
	fmt.Println(string(b))

	i, _ := json.Marshal(1)
	fmt.Println(string(i))

	f, _ := json.Marshal(2.34)
	fmt.Println(string(f))

	s, _ := json.Marshal("gopher")
	fmt.Println(string(s))

	srcslice := []string{"apple", "pear", "peach"}
	dstslice, _ := json.Marshal(srcslice)
	fmt.Println(string(dstslice))

	srcmap := map[string]int{"apple": 5, "lettuce": 7}
	dstmap, _ := json.Marshal(srcmap)
	fmt.Println(string(dstmap))


	srcrsp1 := &RSP1{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	dstrsp1, _ := json.Marshal(srcrsp1)
	fmt.Println(string(dstrsp1))

	srcrsp2 := &RSP2{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	dstrsp2, _ := json.Marshal(srcrsp2)
	fmt.Println(string(dstrsp2))


	byt := []byte(`{"num": 6.13, "str": ["A", "b"]}`)

	var data map[string]interface{}

	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}
	fmt.Println(data)

	num := data["num"].(float64)
	fmt.Println(num)

	strs := data["str"].([]interface{})
	fmt.Println(strs)
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page":1, "fruits": ["apple", "pear"]}`
	res := &RSP2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "pear": 7}
	enc.Encode(d)


}
