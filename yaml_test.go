package yaml

import (
	"fmt"
	"testing"
	"time"
)

func TestFilePath(t *testing.T) {
	New(FilePath("./conf/app.yaml"))
}

type User struct {
	Name string
	Age  int `json:"age"`
	In   struct {
		A string `json:"a"`
		B struct {
			C string `json:"c"`
			D string
		}
	}
	Li  []int          `json:"li"`
	Mp  map[string]int `json:"mp"`
	B   bool           `json:"b"`
	Obj struct {
		A int
		B int
	}
	Date time.Time `json:"date"`
	List [][]int   `json:"list"`
}

func TestConfig_Get(t *testing.T) {
	var u User
	err := New().Get(&u)
	fmt.Println(err)
	fmt.Println(u)
}
