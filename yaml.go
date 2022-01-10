package yaml

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"sync"
)

var (
	cfg  *config
	once sync.Once
)

type Yml interface {
	load() *config
	Get(dest interface{}) error
	String(param ...string) error
}

type config struct {
	path     string
	fileName string
	content  []byte
}

// New returns a structure pointer of YML
func New() Yml {
	once.Do(func() {
		cfg = new(config).load()
	})
	return cfg
}

// load read file information
// The file under conf will be read by default
func (c *config) load() *config {
	if c.path == "" {
		c.path = "./conf/application.yaml"
	}
	content, err := ioutil.ReadFile(c.path)
	if err != nil {
		log.Fatalf("解析config.yaml读取错误: %v", err)
	}
	c.content = content
	return c
}

// Get get the top parameter structure information according to the name
// For example:
//		var u User
// 		c.Get(&u)
func (c *config) Get(dest interface{}) error {

	cp := make(map[string]interface{})
	if err := yaml.Unmarshal(c.content, &cp); err != nil {
		return errors.Wrap(err, "yaml get filed")
	}

	d := &decode{}
	if err := d.unmarshal(dest, cp); err != nil {
		return errors.Wrap(err, "yaml get filed")
	}
	return nil
}

// String gets the value of the specified field
func (c *config) String(param ...string) error {
	return nil
}
