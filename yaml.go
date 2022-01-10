package yaml

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

var (
	cfg  *config
	once sync.Once
)

type Yml interface {
	load() *config
	G(dest interface{}, param ...string) error
	Value(param ...string) (interface{}, error)
}

type config struct {
	dir     string
	name    string
	content []byte
	cp      map[string]interface{}
	d       *decode
}

type Options struct {
	f func(*config)
}

// New returns a structure pointer of YML
func New(options ...Options) Yml {
	once.Do(func() {
		c := &config{
			d: &decode{},
		}
		if options != nil {
			for _, option := range options {
				option.f(c)
			}
		}
		cfg = c.load()
	})
	return cfg
}

func Dir(path string) Options {
	return Options{func(do *config) {
		do.dir = path
	}}
}

func Name(file string) Options {
	return Options{func(do *config) {
		do.name = file
	}}
}

func FilePath(file string) Options {
	return Options{func(do *config) {

		last := strings.LastIndex(file, "/")

		path := file[0 : last+1]
		name := file[last+1:]

		do.dir = path
		do.name = name
	}}
}

// load read file information
// The file under conf will be read by default
func (c *config) load() *config {

	if c.dir == "" {
		c.dir = "./configs/"
	}

	if c.dir[len(c.dir)-1:] != "/" {
		c.dir = c.dir + "/"
	}

	if c.name == "" {
		c.name = "app.yaml"
	}

	content, err := ioutil.ReadFile(c.dir + c.name)
	if err != nil {
		log.Fatalf("解析%s读取错误: %v", c.name, err)
	}
	c.content = content

	cp := make(map[string]interface{})
	if err := yaml.Unmarshal(c.content, &cp); err != nil {
		log.Fatalf("解析%s读取错误: %v", c.name, err)
	}
	c.cp = cp
	return c
}

// G query related structure data
// For example:
//		var u User
// 		c.Get(&u)
func (c *config) G(dest interface{}, param ...string) error {
	return c.d.unmarshal(dest, c.cp, param...)
}

// Value gets the value of the specified field
func (c *config) Value(param ...string) (interface{}, error) {
	var val interface{}
	if err := c.d.unmarshal(&val, c.cp, param...); err != nil {
		return nil, errors.Wrap(err, "string filed")
	}
	return val, nil
}
