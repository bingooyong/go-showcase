package config

import (
	"testing"
	"os"
	"io/ioutil"
)

var text = `
server:
  port: 8080
appname: gopads-server
`

func TestLoad(t *testing.T) {
	tmp := os.TempDir()
	file := tmp + "/tmp.yaml"
	ioutil.WriteFile(file, []byte(text), os.ModePerm)
	defer os.Remove(file)

	config.Load(file)
	t.Log(config)

}
