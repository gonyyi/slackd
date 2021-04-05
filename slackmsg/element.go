package slackmsg

type element struct {
	Type        string `json:"type,omitempty"`
	Text        string `json:"text,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
	AltText     string `json:"alt_text,omitempty"`
	Placeholder *text  `json:"placeholder,omitempty"`
	ActionID    string `json:"action_id,omitempty"`
	Options     []struct {
		Text  text   `json:"text,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"options,omitempty"`
	Elements []element `json:"elements,omitempty"`
}

func (e *element) Reset() {
	e.Type = ""
	e.Text = ""
	e.ImageURL = ""
	e.AltText = ""
	e.Placeholder = nil
	e.ActionID = ""
	e.Options = e.Options[:0]
	e.Elements = e.Elements[:0]
}

func (e *element) AppendText(dst []string) []string {
	if e.Text != "" {
		return append(dst, e.Text)
	}
	for _, v := range e.Elements {
		dst = v.AppendText(dst)
	}
	return dst
}

func (e *element) AsImage(imageURL, altText string) {
	e.Type = "image"
	e.ImageURL = imageURL
	e.AltText = altText
}
func (e *element) AsMarkdown(markdown string) {
	e.Type = "mrkdwn"
	e.Text = markdown
}
