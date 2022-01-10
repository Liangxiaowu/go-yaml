package yaml

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

type decode struct {
	t reflect.Type
	v reflect.Value
}

// unmarshal decodes the document found within the in byte slice
func (d *decode) unmarshal(dest interface{}, cp map[string]interface{}, params ...string) error {
	d.t = reflect.TypeOf(dest).Elem()

	d.v = reflect.ValueOf(dest).Elem()

	if params == nil {
		c, ok := cp[strings.ToLower(d.t.Name())]
		if !ok {
			return errors.Errorf("No related parameters found:%s ", strings.ToLower(d.t.Name()))
		}
		return d.sUnmarshal(c.(map[string]interface{}))
	}

	var (
		c  interface{}
		ok bool
	)
	for _, param := range params {
		c, ok = cp[param]
		if !ok {
			return errors.Errorf("No related parameters found:%s ", param)
		}
		if reflect.TypeOf(c).Kind() == reflect.Map {
			cp = c.(map[string]interface{})
		}
	}
	if reflect.ValueOf(c).Type() == d.t {
		d.v.Set(reflect.ValueOf(c))
		return nil
	}

	k := d.v.Type().Kind()
	switch k {
	case reflect.Interface:
		d.v.Set(reflect.ValueOf(c))
	case reflect.Struct:
		return d.sUnmarshal(cp)
	case reflect.Slice:
		d.deslice(d.v, c.([]interface{}))
		return nil
	case reflect.Map:
		d.demap(d.v, c.(map[string]interface{}))
		return nil
	}

	return nil
}

func (d *decode) sUnmarshal(cp map[string]interface{}) error {
	if d.t.Kind() != reflect.Struct {
		return errors.Errorf("No related structures found:%s ", strings.ToLower(d.t.Name()))
	}

	b := cp
	for i := 0; i < d.t.NumField(); i++ {
		field := d.t.Field(i)

		name := d.getTagName(field)

		if val, ok := b[name]; ok {
			f := d.v.Field(i)
			if reflect.ValueOf(val).Type() == field.Type {
				d.v.Field(i).Set(reflect.ValueOf(val))
				continue
			}

			k := f.Type().Kind()

			switch k {
			case reflect.Struct:
				d.destruct(f, val)
				continue
			case reflect.Slice:
				d.deslice(f, val.([]interface{}))
				continue
			case reflect.Map:
				d.demap(f, val.(map[string]interface{}))
				continue
			}
		}
	}
	return nil
}

func (d *decode) uu(dest interface{}, cp map[string]interface{}, params ...string) error {
	//t := reflect.TypeOf(dest)
	v := reflect.ValueOf(dest).Elem()

	var (
		c  interface{}
		ok bool
	)
	for _, param := range params {
		c, ok = cp[param]
		if !ok {
			return errors.Errorf("No related parameters found:%s ", param)
		}
		if reflect.TypeOf(c).Kind() == reflect.Map {
			cp = c.(map[string]interface{})
		}
	}
	//fmt.Println(reflect.ValueOf(c).Type())
	//fmt.Println(d.v.Type())
	//d.v.Set(reflect.ValueOf(c))
	if reflect.ValueOf(c).Type() == v.Type() {
		//fmt.Println(c)
		//fmt.Println(1111111)
		v.Set(reflect.ValueOf(c))
		return nil
	}

	return nil
}

func (d *decode) getTagName(value reflect.StructField) string {
	name := value.Tag.Get("json")
	if value.Tag.Get("json") == "" {
		name = strings.ToLower(value.Name)
	}
	return name
}

// destruct struct custom type conversion
func (d *decode) destruct(value reflect.Value, mp interface{}) {
	m := mp.(map[string]interface{})
	structField := value.Type()
	for i := 0; i < structField.NumField(); i++ {
		v := value.Field(i)
		name := d.getTagName(structField.Field(i))
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
