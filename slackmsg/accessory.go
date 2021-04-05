package slackmsg


type accessory struct {
	Type        string `json:"type,omitempty"`
	Placeholder struct {
		Type  string `json:"type,omitempty"`
		Text  string `json:"text,omitempty"`
		Emoji bool   `json:"emoji,omitempty"`
	} `json:"placeholder,omitempty"`
	Options []struct {
		Text        text `json:"text,omitempty"`
		Description struct {
			Type string `json:"type,omitempty"`
			Text string `json:"text,omitempty"`
		} `json:"description,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"options,omitempty"`
	Text        text   `json:"text,omitempty"`
	Value       string `json:"value,omitempty"`
	URL         string `json:"url,omitempty"`
	ActionID    string `json:"action_id,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
	AltText     string `json:"alt_text,omitempty"`
	InitialTime string `json:"initial_time,omitempty"`
}
func (a *accessory) Reset() {
	a.Type = ""
	a.Placeholder.Type = ""
	a.Placeholder.Text = ""
	a.Placeholder.Emoji = false
	a.Options = a.Options[:0]
	a.Text.Reset()
	a.Value = ""
	a.URL = ""
	a.ActionID = ""
	a.ImageURL = ""
	a.AltText = ""
	a.InitialTime = ""

}