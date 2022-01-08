package yaml

import (
	"fmt"
	"testing"
)

func bil() YamlInter {
	return New()
}

type User struct {
	Name string
	Age  int `json:"age"`
	In   struct {
		A string `json:"a"`
		B string `json:"b"`
	}
	Li []int `json:"li"`
}

func TestConfig_Get(t *testing.T) {
	var u User
	err := bil().Get(&u)
	fmt.Println(err)
	fmt.Println(u)
}
