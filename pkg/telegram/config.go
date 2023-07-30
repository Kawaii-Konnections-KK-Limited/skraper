package telegram

import (
	"bufio"
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Channels []string `yaml:"channels"`
}

func GetConfig(path string, c *Config) {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	r := bufio.NewReader(fi)

	d := yaml.NewDecoder(r)

	d.Decode(c)
}
