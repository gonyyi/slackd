package slackmsg

import "sync"

var pool = sync.Pool{
	New: func() interface{} {
		return &SlackMsg{}
	},
}

func Get(slackToken string, moduleDir string) *SlackMsg {
	s := pool.Get().(*SlackMsg)
	s.moduleDir = moduleDir
	s.slackToken = slackToken
	return s
}

func Put(m *SlackMsg) {
	m.Incoming.Reset()
	m.Outgoing.Reset()
	pool.Put(m)
}

type SlackMsg struct {
	slackToken string
	moduleDir  string
	Incoming   incoming
	Outgoing   outgoing
}
