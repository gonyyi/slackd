package slackmsg


// ====================================================================
// TEXT
// ====================================================================
type text struct {
	Type  string `json:"type,omitempty"`
	Text  string `json:"text,omitempty"`
	Emoji bool   `json:"emoji,omitempty"`
}
func (t *text) Reset() {
	t.Type = ""
	t.Text = ""
	t.Emoji = false
}
