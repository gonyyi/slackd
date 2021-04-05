package slackmsg


// ====================================================================
// ATTACHMENT
// ====================================================================
type attachment struct {
	Fallback   string  `json:"fallback,omitempty"`    // Plain-text summary of the attachment.",
	Color      string  `json:"color,omitempty"`       // #2eb886",
	Pretext    string  `json:"pretext,omitempty"`     // Optional text that appears above the attachment block",
	AuthorName string  `json:"author_name,omitempty"` // Bobby Tables",
	AuthorLink string  `json:"author_link,omitempty"` // http://flickr.com/bobby/",
	AuthorIcon string  `json:"author_icon,omitempty"` // http://flickr.com/icons/bobby.jpg",
	Title      string  `json:"title,omitempty"`       // Slack API Documentation",
	TitleLink  string  `json:"title_link,omitempty"`  // https://api.slack.com/",
	Text       string  `json:"text,omitempty"`        // Optional text that appears within the attachment",
	Fields     []field `json:"fields,omitempty"`
	ImageURL   string  `json:"image_url,omitempty"`   // http://my-website.com/path/to/image.jpg",
	ThumbURL   string  `json:"thumb_url,omitempty"`   // http://example.com/path/to/thumb.png",
	Footer     string  `json:"footer,omitempty"`      // Slack API",
	FooterIcon string  `json:"footer_icon,omitempty"` // https://platform.slack-edge.com/img/default_application_icon.png",
	Ts         int64   `json:"ts,omitempty"`          // 123456789
	Blocks     []block `json:"blocks,omitempty"`
}

func (a *attachment) Reset() {
	a.Fallback = ""
	a.Color = ""
	a.Pretext = ""
	a.AuthorIcon = ""
	a.AuthorLink = ""
	a.AuthorIcon = ""
	a.Title = ""
	a.TitleLink = ""
	a.Text = ""
	a.Fields = a.Fields[:0]
	a.ImageURL = ""
	a.ThumbURL = ""
	a.Footer = ""
	a.FooterIcon = ""
	a.Ts = 0
	a.Blocks = a.Blocks[:0]
}