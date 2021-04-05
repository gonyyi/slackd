package email

import "testing"

func TestSendMail(t *testing.T) {
	a := NewMailer()
	err := a.Send("gonyi@acxiom.com", "myVerificationCode123da")
	if err != nil {
		println(err.Error())
	}
}
