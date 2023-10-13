package study

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func TestJson() {
	// 1. 序列化(将基本数据类型转换为json字符串)
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	floatB, _ := json.Marshal(2.34)
	fmt.Println(string(floatB))

	stringB, _ := json.Marshal("gopher")
	fmt.Println(string(stringB))

	slcB, _ := json.Marshal([]string{"a", "b", "c"})
	fmt.Println(string(slcB))

	mapB, _ := json.Marshal(map[string]int{"a": 1, "b": 2})
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// 2. 反序列化(将json字符串转换为基本数据类型)
	var jsonBlob = []byte(`{"page": 1, "fruits": ["apple", "peach"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(jsonBlob, &dat); err != nil {
		panic(err)
	}
	fmt.Println("dat:", dat)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println("res:", res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)

	d := map[string]int{"apple": 1, "pear": 2}

	enc.Encode(d)
}
