package slackmsg

type block struct {
	Type      string     `json:"type,omitempty"`
	BlockID   string     `json:"block_id,omitempty"` // "D7US",
	Text      *text      `json:"text,omitempty"`
	Fields    []text     `json:"fields,omitempty"`
	Accessory *accessory `json:"accessory,omitempty"`
	Elements  []element  `json:"elements,omitempty"`
	Title     *text      `json:"title,omitempty"`
	ImageURL  string     `json:"image_url,omitempty"`
	AltText   string     `json:"alt_text,omitempty"`
}

func (b *block) Reset() {
	b.Type = ""
	b.BlockID = ""
	b.Text = nil
	b.Fields = b.Fields[:0]
	b.Accessory = nil
	b.Elements = b.Elements[:0]
	b.Title = nil
	b.ImageURL = ""
	b.AltText = ""
}

func (b *block) AppendText(dst []string) []string {
	for _, v := range b.Elements {
		dst = v.AppendText(dst)
	}
	return dst
}
