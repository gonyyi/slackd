package slackmsg

type customIncomingType string
type customOutgoingType string

const (
	MSG_NEW    customIncomingType = "new"
	MSG_UPDATE customIncomingType = "update"
	MSG_DELETE customIncomingType = "delete"

	MSG_INFO     customOutgoingType = "info"
	MSG_ERROR    customOutgoingType = "error"
	MSG_WARNING  customOutgoingType = "warning"
	MSG_MARKDOWN customOutgoingType = "markdown"
)

type customIncoming struct {
	Token        string             `json:"token,omitempty"`
	EventChannel string             `json:"eventChannel,omitempty"`
	EventTS      string             `json:"eventTS,omitempty"`
	Text         string             `json:"text,omitempty"`
	Type         customIncomingType `json:"type,omitempty"` // new, update, delete
	User         string             `json:"userID,omitempty"`
}

func (ci *customIncoming) Reset() {
	ci.Token = ""
	ci.EventChannel = ""
	ci.EventTS = ""
	ci.Text = ""
	ci.Type = ""
	ci.User = ""
}

type customOutgoing struct {
	Token        string             `json:"token,omitempty"`
	EventChannel string             `json:"eventChannel,omitempty"`
	EventTS      string             `json:"eventTS,omitempty"`
	Messages      []customOutgoingMsg  `json:"messages,omitempty"`
	ReplyInThread bool                 `json:"replyInThread,omitempty"`
	Files         []customOutgoingFile `json:"files,omitempty"`
}

func (co *customOutgoing) Reset() {
	co.Token = ""
	co.EventChannel = ""
	co.EventTS = ""
	co.Messages = co.Messages[:0]
	co.ReplyInThread = false
	co.Files = co.Files[:0]
}

type customOutgoingMsg struct {
	Text string             `json:"text,omitempty"`
	Type customOutgoingType `json:"type,omitempty"` // info, error, warning, code snippet, markdown //(default: info)
}

func (m *customOutgoingMsg) Reset() {
	m.Text = ""
	m.Type = ""
}

type customOutgoingFile struct {
	Filename string `json:"filename,omitempty"`
	Title    string `json:"title,omitempty"`
	Comment  string `json:"comment,omitempty"`
	Gzip     bool   `json:"gzip,omitempty"`
}

func (f customOutgoingFile) Upload(slackToken, channel, threadTS, dir string) error {
	return postFile(slackToken, channel, threadTS, dir, f.Filename, f.Title, f.Comment, f.Gzip)
}
