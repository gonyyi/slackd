package modules

import (
	"fmt"
	"testing"
)

func TestNewModules(t *testing.T) {
	m := NewModules("/Users/gonyi/go/src/github.com/gonyyi/slackd/modules")
	m.Scan()

	f := func(s string) {
		// path, pCmd, pArgs, availableCmd string, ok bool
		c := m.ParseCommand(s)
		fmt.Printf("Cmd:  <%s>\n  Path: %s\n  pCmd: %s\n  Args: %s\n  aCmd: %s\n  isMo: %t\n  isPa: %t\n  isHelp: %t\n  isSyst: %t\n  intf-v: %d\n  Erro  : %s\n",
			s, c.ModulePath, c.ParsedCmd, c.ParsedArgs, c.AvailableCmd("|"), c.IsModule(), c.IsParsed(), c.IsHelp(), c.IsSystem(), c.InterfaceVersion(), c.Error())
		if c.module != nil {
			println("  Module.Name: "+c.module.Info.Name)
		}
		println()
		{
			sout, serr, err := c.module.Exec([]byte(`{"custom":{"incoming":{"user":"gon"}}}`))
			if err != nil {println("err:", err.Error())}
			println("sout:", string(sout))
			println("serr:", string(serr))
		}
	}

	test := []string{
		"testModBasic abcdef 123",
	}

	f(test[0])
}
