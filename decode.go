package yaml

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

type decode struct {
	errs []error
}

// unmarshal decodes the document found within the in byte slice
func (d *decode) unmarshal(dest interface{}, cp interface{}) error {

	c := cp.(map[string]interface{})

	v := reflect.TypeOf(dest).Elem()

	body, ok := c[strings.ToLower(v.Name())]
	if !ok {
		return errors.Errorf("No related parameters found:%s ", strings.ToLower(v.Name()))
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
			k := s.Field(i).Type().Kind()
			//
			if reflect.ValueOf(val).Type() == field.Type {
				s.Field(i).Set(reflect.ValueOf(val))
				continue
			}

			switch k {
			case reflect.Struct:
				d.destruct(s.Field(i), val)
				continue
			case reflect.Slice:
				d.deslice(s.Field(i), val.([]interface{}))
				continue
			case reflect.Map:
				d.demap(s.Field(i), val.(map[string]interface{}))
				continue
			}
		}
	}
	return nil
}

// destruct struct custom type conversion
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

// deslice slice custom type conversion
func (d *decode) deslice(value reflect.Value, li []interface{}) {
	fmt.Println(value.Type().Elem().String())
	switch value.Type().String() {
	case "[]int":
		var list []int
		for _, l := range li {
			if reflect.TypeOf(l).String() == "int" {
				list = append(list, l.(int))
			}
		}
		value.Set(reflect.ValueOf(list))
	case "[]string":
		var list []string
		for _, l := range li {
			if reflect.TypeOf(l).String() == "string" {
				list = append(list, l.(string))
			}
		}
		value.Set(reflect.ValueOf(list))
	}
}

// deslice map custom type conversion
func (d *decode) demap(value reflect.Value, mp map[string]interface{}) {
	switch value.Type().String() {
	case "map[string]string":
		mm := make(map[string]string)
		for i, m := range mp {
			fmt.Println(i, m)
			fmt.Println(reflect.TypeOf(i).String())
			if reflect.TypeOf(m).String() == "string" {
				mm[i] = m.(string)
			}
		}
		value.Set(reflect.ValueOf(mm))
	case "map[string]int":
		mm := make(map[string]int)
		for i, m := range mp {
			if reflect.TypeOf(m).String() == "int" {
				mm[i] = m.(int)
			}
		}
		value.Set(reflect.ValueOf(mm))
	}
}
