package yaml

import (
	"gopkg.in/yaml.v3"
	"os"
)

func Parse(path string, d any) {
	buf, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(buf, d)
	if err != nil {
		panic(err)
	}
}
