package slackd

import "testing"

func TestNew(t *testing.T) {
	s, err := New("./cmd/shorty_conf.json")
	if err != nil {
		println(err.Error())
		return
	}
	s.Run(":8080")
}