package modules

import (
	"bytes"
	"os/exec"
	"path"
)

// interface_version 3
func (m *module) Exec(jsonData []byte) (jsonOutput []byte, errorReceived []byte, err error) {
	var c *exec.Cmd

	c = exec.Command("./" + m.Module.EntryPoint)
	c.Dir = path.Join(m.dir)
	c.Stdin = bytes.NewBuffer(jsonData)

	var sout bytes.Buffer
	var serr bytes.Buffer
	c.Stdout = &sout
	c.Stderr = &serr

	err = c.Run()

	return sout.Bytes(), serr.Bytes(), err
}

