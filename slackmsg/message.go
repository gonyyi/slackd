package slackmsg

type message struct {
	ClientMsgID string  `json:"client_msg_id,omitempty"`
	Type        string  `json:"type,omitempty"` // "message",
	Text        string  `json:"text,omitempty"` // "hello",
	User        string  `json:"user,omitempty"` // "UFLEM86PP",
	Team        string  `json:"team,omitempty"` // "TCVUCDDDY",
	Edited      edited  `json:"edited,omitempty"`
	Blocks      []block `json:"blocks,omitempty"`
	TS          string  `json:"ts,omitempty"` // "1601562074.000300",
	SourceTeam  string  `json:"source_team,omitempty"`
	UserTeam    string  `json:"user_team,omitempty"`
	ChannelType string  `json:"channel_type,omitempty"`
}

func (m *message) Reset() {
	m.ClientMsgID = ""
	m.Type = ""
	m.Text = ""
	m.User = ""
	m.Team = ""
	m.Edited.Reset()
	m.Blocks = m.Blocks[:0]
	m.TS = ""
	m.SourceTeam = ""
	m.UserTeam = ""
	m.ChannelType = ""
}

func (m *message) AppendText(dst []string) []string {
	if m.Text != "" {
		return append(dst, m.Text)
	}

	for _, v := range m.Blocks {
		dst = v.AppendText(dst)
	}

	return dst
}
