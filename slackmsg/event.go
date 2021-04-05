package slackmsg

type event struct {
	ClientMsgID string  `json:"client_msg_id,omitempty"` // "f5f5b9e5-69b8-4fb9-af18-5bf8940053f9",
	Type        string  `json:"type,omitempty"`          // "message",
	SubType     string  `json:"subtype,omitempty"`       // "",
	Text        string  `json:"text,omitempty"`          // "hello",
	User        string  `json:"user,omitempty"`          // "UFLEM86PP",
	TS          string  `json:"ts,omitempty"`            // "1601562074.000300",
	EventTS     string  `json:"event_ts,omitempty"`      // "1601562074.000300",
	Hidden      bool    `json:"hidden,omitempty"`
	DeletedTS   string  `json:"deleted_ts,omitempty"`
	Team        string  `json:"team,omitempty"` // "TCVUCDDDY",
	Blocks      []block `json:"blocks,omitempty"`
	Message     message `json:"message,omitempty"`
	PrevMessage message `json:"previous_message,omitempty"`
	Channel     string  `json:"channel,omitempty"`      // "D015QGYPV0F",
	ChannelType string  `json:"channel_type,omitempty"` // "im"
	BotID       string  `json:"bot_id,omitempty"`
}

func (e *event) Reset() {
	e.ClientMsgID = ""
	e.Type = ""
	e.SubType = ""
	e.Text = ""
	e.User = ""
	e.TS = ""
	e.EventTS = ""
	e.Hidden = false
	e.DeletedTS = ""
	e.Team = ""
	e.Blocks = e.Blocks[:0]
	e.Message.Reset()
	e.PrevMessage.Reset()
	e.Channel = ""
	e.ChannelType = ""
	e.BotID = ""
}
