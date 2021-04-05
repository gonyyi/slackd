package slackd

import (
	"github.com/gonyyi/slackd/slackmsg"
	"io/ioutil"
	"net/http"
)

func theHandler(s *Slackd) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.log.Error(HTTP).Int("status", 400).Err("err", err).Write(ERRS_FAILED_READ_REQ_BODY)
			w.WriteHeader(400)
			return
		}

		// PROCESS MESSAGE
		{
			m := slackmsg.Get(s.conf.Slack.Token, s.conf.Modules.Dir)
			defer slackmsg.Put(m)

			if err := m.Incoming.Load(body); err != nil {
				s.log.Error(HTTP).
					Err("err", err).
					Str("uri", r.RequestURI).
					Int("msg.size", len(body)).
					Write(ERRS_FAILED_LOAD_INCOMING_MSG)
				return
			}

			if code, ok := m.Incoming.IsChallengeMsg(); ok {
				w.Write([]byte(code))
				s.log.Info(HTTP).
					Str("challenge", code).
					Write(INFO_RECEIVED_SLACK_CHALLENGE)
				return
			}

			s.log.Trace(REQ).
				Str("userID", m.Incoming.Custom.User).
				Str("command", m.Incoming.Custom.Text).
				Write("")

			c := s.modules.ParseCommand(m.Incoming.Custom.Text)
			switch {
			case c.IsSystem(): // system command
				switch c.ParsedCmd {
				case "info":
					// show system info
				case "i am your daddy":
					// check if this person can be admin
				default:
					// need admin access
				}
			case c.IsHelp(): // help command
				if c.IsParsed() {
					// Return help menu
					// c.Module().Help.Intro
				} else {
					// request command not found from help menu
				}
			case c.IsParsed() && !c.IsModule(): // parsed but not the module
				// c.AvailableCmd(", ")
				//
			case c.IsParsed() && c.IsModule(): // parsed module; run the module here

			default:
				// unrecognized command
			}
		}
		s.log.Warn(SYS).Write(WARNS_UNEXPECTED_EXCEPTION)
		return
	}
}
