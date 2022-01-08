package yaml

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"sync"
)

var (
	cfg  *config
	once sync.Once
)

type YamlInter interface {
	load() *config
	Get(dest interface{}) error
	String(param ...string) error
}

type config struct {
	path     string
	fileName string
	content  []byte
}

func New() YamlInter {
	once.Do(func() {
		cfg = new(config).load()
	})
	return cfg
}

func (c *config) load() *config {
	if c.path == "" {
		c.path = "./conf/test.yaml"
	}
	content, err := ioutil.ReadFile(c.path)
	if err != nil {
		log.Fatalf("解析config.yaml读取错误: %v", err)
	}
	c.content = content
	return c
}

func (c *config) Get(dest interface{}) error {
	cp := make(map[string]interface{})
	if err := yaml.Unmarshal(c.content, &cp); err != nil {

	}

	d := &decode{}
	d.unmarshal(dest, cp)
	return nil
}

func (c *config) String(param ...string) error {
	return nil
}
