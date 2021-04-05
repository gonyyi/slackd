package slackmsg

import (
	"encoding/json"
	"github.com/gonyyi/atype"
)

const (
	ERR_BOT_MSG       atype.ErrorStr = "bot message"
	ERR_CHALLENGE_MSG atype.ErrorStr = "slack challenge message"
)

func printJSON(j interface{}) {
	b, err := json.MarshalIndent(j, "", "   ")
	if err != nil {
		println("printJSON: err=" + err.Error())
		return
	}
	println(string(b))
}

func checkResp(b []byte) (err error) {
	var a struct {
		Ok    bool   `json:"ok"`
		Error string `json:"error"`
	}
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	if a.Ok {
		return nil
	}
	if a.Error == "" {
		return atype.ErrorStr("unknown error")
	}
	return atype.ErrorStr(a.Error)
}
