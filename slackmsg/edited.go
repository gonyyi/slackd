package slackmsg

type edited struct {
	User string `json:"user,omitempty"` // "UFLEM86PP",
	TS   string `json:"ts,omitempty"`   // "1601562074.000300",
}

func (e *edited) Reset() {
	e.User = ""
	e.TS = ""
}
