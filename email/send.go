package email

import (
	"fmt"
	"os/exec"
	"strings"
)

func NewMailer() *Mail {
	m := Mail{
		MailCmdPath:  "/usr/bin/mail",
		Subject:      "Slackd verification code",
		BodyTemplate: "Your verification code is <$CODE$>.\nTo activate, go to slack and direct message to <Slackd> `i am $CODE$` to verify.",
	}
	m.getMailCmdPath()
	return &m
}

type Mail struct {
	MailCmdPath  string
	Subject      string
	BodyTemplate string
}

func (m *Mail) getMailCmdPath() error {
	s, err := exec.LookPath("mail")
	if err != nil {
		return err
	}
	m.MailCmdPath = s
	return nil
}

func (m *Mail) Send(to, code string) error {
	c := exec.Command(m.MailCmdPath, "-s", m.Subject, to)
	stdin, err := c.StdinPipe()
	if err != nil {
		return fmt.Errorf("%s: %w", "StdinPipe error", err)
	}

	stdin.Write([]byte(strings.ReplaceAll(m.BodyTemplate, "$CODE$", code)))
	stdin.Close()
	return c.Run()
}
