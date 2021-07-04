package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type resp1 struct {
	Page   int
	Fruits []string
}

type resp2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

var log = fmt.Println

func main() {
	bolB, _ := json.Marshal(true)

	log("bool", string(bolB))

	slcD := []string{"apple", "peach", "pear"}

	slcB, _ := json.Marshal(slcD)

	log(slcD, string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	log(mapD, string(mapB))

	// resp1D := resp1{
	resp1D := &resp1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	resp1B, _ := json.Marshal(resp1D)
	log(resp1D, string(resp1B))

	resp2D := &resp2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	resp2B, _ := json.Marshal(resp2D)
	log(resp2D, string(resp2B))

	// decoding JSON data into Go values
	byt := []byte(`{"num": 6.13, "strs": ["apple", "peach", "pear"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	log(dat)

	log(dat["num"].(float64))

	// log(dat["strs"].([]interface{})[0].(string))
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	log(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	// var res resp2
	res := resp2{}
	json.Unmarshal([]byte(str), &res)

	log(res)
	log(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode((d))
}
