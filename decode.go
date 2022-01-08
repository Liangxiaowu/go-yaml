package yaml

import (
	"fmt"
	"reflect"
	"strings"
)

type decode struct {
}

func (d *decode) unmarshal(dest interface{}, cp interface{}) error {
	// 获取结构体实例的反射类型对象
	c := cp.(map[string]interface{})
	fmt.Println(c)
	v := reflect.TypeOf(dest).Elem()
	body, ok := c[strings.ToLower(v.Name())]
	if !ok {
		// 查不到值
	}
	s := reflect.ValueOf(dest).Elem()
	b := body.(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		name := field.Tag.Get("json")
		if name == "" {
			name = strings.ToLower(field.Name)
		}

		if val, ok := b[name]; ok {
			if reflect.ValueOf(val).Type() == field.Type {
				s.Field(i).Set(reflect.ValueOf(val))
				continue
			}

			if s.Field(i).Type().Kind() == reflect.Struct {
				d.destruct(s.Field(i), val)
				continue
			}

		}

	}

	return nil
}

// destruct
func (d *decode) destruct(value reflect.Value, mp interface{}) {
	m := mp.(map[string]interface{})
	structField := value.Type()
	for i := 0; i < structField.NumField(); i++ {
		v := value.Field(i)
		name := structField.Field(i).Tag.Get("json")
		if structField.Field(i).Tag.Get("json") == "" {
			name = strings.ToLower(structField.Field(i).Name)
		}
		if val, ok := m[name]; ok {
			// 判断是否还是结构体
			if v.Type().Kind() == reflect.Struct {
				d.destruct(value.Field(i), val)
				continue
			}
			v.Set(reflect.ValueOf(val))
		}
	}
}
