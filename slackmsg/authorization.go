package slackmsg

type authorization struct {
	EnterpriseID        string `json:"enterprise_id,omitempty"`
	TeamID              string `json:"team_id,omitempty"`
	UserID              string `json:"user_id,omitempty"`
	IsBot               bool   `json:"is_bot,omitempty"`
	IsEnterpriseInstall bool   `json:"is_enterprise_install,omitempty"`
}
