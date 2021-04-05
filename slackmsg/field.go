package slackmsg



// ====================================================================
// FIELD
// ====================================================================
type field struct {
	Title string `json:"title,omitempty"` // Priority",
	Value string `json:"value,omitempty"` // High",
	Short bool   `json:"short,omitempty"` // false
}
func (f *field) Reset() {
	f.Title = ""
	f.Value = ""
	f.Short = false
}
