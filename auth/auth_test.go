package auth

import (
	"github.com/gonyyi/mutt"
	"testing"
)

func Test1(t *testing.T) {
	b := []byte(`{"id":"123", "tokens":{"gon":"a"}}`)
	u, _ := mutt.NewUser(b)
	println(u.Created.String())

	u.LastLogin.Set()
	u.LastModified = 0
	println(u.LastModified.String())
	out, _ := u.Bytes()
	println(string(out))
}
