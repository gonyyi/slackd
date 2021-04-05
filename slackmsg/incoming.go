package slackmsg

import (
	"encoding/json"
	"strings"
)

type incoming struct {
	Token              string          `json:"token,omitempty"`      // "etlXATMDxTR3iWNXir47ksC7"
	TeamID             string          `json:"team_id,omitempty"`    // "TCVUCDDDY",
	ApiAppID           string          `json:"api_app_id,omitempty"` // "AGTCUQE0J",
	Event              event           `json:"event,omitempty"`
	Type               string          `json:"type,omitempty"`       // "event_callback",
	Challenge          string          `json:"challenge,omitempty"`  // this is only for slack API's challenge
	EventID            string          `json:"event_id,omitempty"`   // "Ev01BBQ6GCMD",
	EventTime          int64           `json:"event_time,omitempty"` // 1601562074,
	Authorizations     []authorization `json:"authorizations,omitempty"`
	IsExtSharedChannel bool            `json:"is_ext_shared_channel,omitempty"`
	EventContext       string          `json:"event_context,omitempty"` // "1-message-TCVUCDDDY-D015QGYPV0F"
	Custom             customIncoming  `json:"Custom,omitempty"`
}

func (inc *incoming) Load(b []byte) (err error) {
	inc.Reset()
	if err = json.Unmarshal(b, inc); err != nil {
		return err
	}
	if _, ok := inc.IsChallengeMsg(); ok {
		return ERR_CHALLENGE_MSG
	}
	if inc.IsBotMsg() {
		return ERR_BOT_MSG
	}
	inc.process()
	return nil
}

func (inc *incoming) IsChallengeMsg() (challegeCode string, ok bool) {
	if inc.Type == "url_verification" && inc.Challenge != "" {
		return inc.Challenge, true
	}
	return "", false
}

func (inc *incoming) IsBotMsg() bool {
	if inc.Event.BotID == "" && inc.Event.SubType != "file_share" && inc.Event.SubType != "bot_message" {
		return false
	}
	return true
}

func (inc *incoming) IsDeleteMsg() bool {
	if inc.Event.SubType == "message_deleted" {
		return true
	}
	return false
}

func (inc *incoming) Reset() {
	inc.Token = ""
	inc.TeamID = ""
	inc.ApiAppID = ""
	inc.Event.Reset()
	inc.Type = ""
	inc.Challenge = ""
	inc.EventID = ""
	inc.EventTime = 0
	inc.Authorizations = inc.Authorizations[:0]
	inc.IsExtSharedChannel = false
	inc.EventContext = ""
}

func (inc *incoming) process() {
	// Get incoming messages to text
	{
		if inc.Event.Text != "" {
			inc.Custom.Text = inc.Event.Text
		} else {
			var text []string
			text = inc.Event.Message.AppendText(text)
			text = inc.Event.PrevMessage.AppendText(text)

			for _, v := range inc.Event.Blocks {
				text = v.AppendText(text)
			}

			inc.Custom.Text = strings.Join(text, " ")
		}
	}

	// Get userID
	{
		if inc.Event.User != "" {
			inc.Custom.User = inc.Event.User
		} else if inc.Event.Message.User != "" {
			inc.Custom.User = inc.Event.Message.User
		} else if inc.Event.PrevMessage.User != "" {
			inc.Custom.User = inc.Event.PrevMessage.User
		}
	}

	// Get incoming message type
	if inc.Event.Type == "message" {
		switch inc.Event.SubType {
		case "message_deleted":
			inc.Custom.Type = MSG_DELETE
		case "message_changed":
			inc.Custom.Type = MSG_UPDATE
		default:
			inc.Custom.Type = MSG_NEW
		}
	}

	inc.Custom.EventChannel = inc.Event.Channel
	inc.Custom.EventTS = inc.Event.TS
}
