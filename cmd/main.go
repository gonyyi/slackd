package main

import (
	"github.com/gonyyi/slackd"
)

func main() {
	test()
}

func t1() {
	sl, err := slackd.New("./cmd/shorty_conf.json")
	if err != nil {
		println(err.Error())
	}

	sl.Run("")
}

func test() {

}
