package slackmsg

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// Msg outgoing: MODULE ==> USER

type outgoing struct {
	Token       string         `json:"token"`               // "xoxb-437621666130-136vHKVZuCqFYexehzNvNnBm",
	Channel     string         `json:"channel"`             // "D015QGYP",
	ThreadTs    string         `json:"thread_ts,omitempty"` // "16015620.000300"
	Text        string         `json:"text,omitempty"`      // "did you just say `hello`",
	Attachments []attachment   `json:"attachments,omitempty"`
	Blocks      []block        `json:"blocks,omitempty"`
	Custom      customOutgoing `json:"Custom,omitempty"`
	dir         string
}

func (o *outgoing) Load(b []byte) error {
	o.Reset()
	if err := json.Unmarshal(b, o); err != nil {
		return err
	}
	return nil
}

func (o *outgoing) Set(token, dir string) {
	o.Token = token
	o.dir = dir
}

func (o *outgoing) Reset() {
	o.Token = ""
	o.Channel = ""
	o.ThreadTs = ""
	o.Text = ""
	o.Attachments = o.Attachments[:0]
	o.Blocks = o.Blocks[:0]
	o.Custom.Reset()
}

// process will take Custom and update as its necessity
func (o *outgoing) Process() error {
	// Update message
	for _, m := range o.Custom.Messages {
		switch m.Type {
		case MSG_INFO:
			o.AddInfo(m.Text)
		case MSG_ERROR:
			o.AddError(m.Text)
		case MSG_WARNING:
			o.AddWarning(m.Text)
		case MSG_MARKDOWN:
			o.AddMarkdown(m.Text)
		}
	}
	if o.Custom.ReplyInThread == false {
		o.ThreadTs = ""
	}

	// Update files
	if len(o.Custom.Files) > 0 {
		var tmpThreadTS, tmpChannel, tmpToken = o.ThreadTs, o.Channel, o.Token

		if o.Custom.EventTS != "" {
			tmpThreadTS = o.Custom.EventTS
		}
		if o.Custom.EventChannel != "" {
			tmpChannel = o.Custom.EventChannel
		}
		if o.Custom.Token != "" {
			tmpToken = o.Custom.Token
		}

		if !o.Custom.ReplyInThread {
			tmpThreadTS = ""
		}

		for _, f := range o.Custom.Files {
			if err := f.Upload(tmpToken, tmpChannel, tmpThreadTS, o.dir); err != nil {
				o.AddError("Uploading `" + f.Filename + "` failed. (Error: `" + err.Error() + "`)")
			}
		}
	}

	o.Custom.Reset()
	return nil
}

// Post will post the message to slack
func (o *outgoing) Post() error {
	// pl, err := json.MarshalIndent(o, "", "  ")
	pl, err := json.Marshal(o)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", SLACK_ENDPOINT_MSG, bytes.NewBuffer(pl))
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+o.Token)

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if _, err = io.Copy(&buf, resp.Body); err != nil {
		return err
	}

	return checkResp(buf.Bytes())
}

func (o *outgoing) AddWarning(markdown string) {
	// #F4D03F
	o.addAttachment("WARNING", "#F4D03F", markdown)
}

func (o *outgoing) AddError(markdown string) {
	// #E74C3C
	o.addAttachment("ERROR", "#E74C3C", markdown)
}

func (o *outgoing) AddInfo(markdown string) {
	// #85C1E9
	o.addAttachment("INFO", "#85C1E9", markdown)
}

func (o *outgoing) addAttachment(name, color, markdown string) {
	o.Attachments = append(o.Attachments, attachment{
		Color: color,
		Blocks: []block{
			{
				Type: "context",
				Elements: []element{
					{
						Type: "mrkdwn",
						Text: "*" + name + "*",
					},
					{
						Type: "mrkdwn",
						Text: markdown,
					},
				},
			},
		},
	})
}

func (o *outgoing) AddMarkdown(markdown string) {
	o.Blocks = append(o.Blocks, block{
		Type: "section",
		Text: &text{
			Type: "mrkdwn",
			Text: markdown,
		},
	})
}

func (o *outgoing) AddDivider() {
	o.Blocks = append(o.Blocks, block{Type: "divider"})
}

func (o *outgoing) AddContext(msgEle ...element) {
	o.Blocks = append(o.Blocks, block{
		Type:     "context",
		Elements: msgEle,
	})
}
