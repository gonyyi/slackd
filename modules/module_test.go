package modules

import (
	"encoding/json"
	"path"
	"testing"
)

func TestModule_Load(t *testing.T) {
	p := "/Users/gonyi/go/src/github.com/gonyyi/slackd/module/testModule/module.json"

	testin := `{"custom":{"incoming":{"user":"gon"}}}`

	println("dir:", path.Dir(p))
	println("file:", path.Base(p))

	m, err := newModule(path.Dir(p))
	if err != nil {println(err.Error());return}

	if false {
		out, _ := json.MarshalIndent(m, "", "   ")
		println(string(out))
	}

	out, _, _ := m.Exec([]byte(testin))
	println(string(out))
}
