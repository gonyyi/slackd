package main

import (
	"errors"
	"github.com/gonyyi/slackd"
)

func main() {
	sl, _ := slackd.New("")

	err := sl.Check2()

	if err != nil {
		println(err.Error())
		println(errors.Is(err, slackd.ERR_A))
		println(errors.Is(err, slackd.ERR_B))
		println(err == slackd.ERR_C)
	} else {
		println("err is nil")
	}
}
