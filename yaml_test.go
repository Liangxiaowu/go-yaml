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

type Obj struct {
	A int
	B int
}

func TestConfig_G(t *testing.T) {
	obj := &Obj{}
	err := New().G(obj, "user", "obj")
	fmt.Println(err)
	fmt.Println(obj)
}

func TestConfig_G2(t *testing.T) {
	user := &User{}
	err := New().G(user)
	fmt.Println(err)
	fmt.Println(user)
}

//func TestConfig_G3(t *testing.T) {
//	var name string
//	err := New().G(name, "user", "name")
//	fmt.Println(err)
//	fmt.Println(name)
//}

func TestConfig_Value(t *testing.T) {
	i, err := New().Value("user", "name")
	fmt.Println(i.(string))
	fmt.Println(err)
}
