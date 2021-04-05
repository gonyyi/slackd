package config

import (
	"encoding/json"
	"io/ioutil"
)

func NewConfig(filename string) (*Config, error) {
	c := Config{}
	if err := c.ReadFile(filename); err != nil {
		return nil, err
	}
	if err := c.validate(); err != nil {
		return nil, err
	}
	return &c, nil
}

type Config struct {
	System struct {
		Logging struct {
			Level     string `json:"level"`
			Filename  string `json:"filename"`
			MaxSizeKB int64  `json:"max_size_kb"`
		} `json:"logging"`
		DataFile string `json:"db_file"`
	} `json:"system"`
	Service struct {
		Name    string   `json:"name"`
		Version string   `json:"version"`
		Host    string   `json:"host"`
		Admin   []string `json:"admin"`
	} `json:"service"`
	Slack struct {
		Token          string `json:"token"`
		Syscmd         string `json:"syscmd"`           // system command shows current status, etc.
		AcceptChallege bool   `json:"accept_challenge"` // this is one time URL verification from slack. (https://api.slack.com/events/url_verification)
	} `json:"slack"`
	Modules struct {
		Dir  string `json:"directory"`
	} `json:"modules"`
}

func (c *Config) ReadFile(filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, &c); err != nil {
		return ERR_BAD_CONF_FILE
	}
	return c.validate()
}

func (c *Config) validate() error {
	if !isOneOf(c.System.Logging.Level, "TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL") {
		return ERR_CONF_LOGGING_INCORRECT_LEVEL
	}
	if c.System.Logging.Filename == "" {
		return ERR_CONF_LOGGING_FILENAME_MISSING
	}
	if c.System.Logging.MaxSizeKB < 1 {
		return ERR_CONF_LOGGING_INVALID_MAXSIZE
	}
	if c.System.DataFile == "" {
		return ERR_CONF_SYSTEM_DB_MISSING
	}
	if c.Service.Name == "" {
		return ERR_CONF_SERVICE_NAME
	}
	if c.Service.Version == "" {
		return ERR_CONF_SERVICE_VERSION
	}
	if c.Service.Host == "" {
		return ERR_CONF_SERVICE_HOST
	}
	if len(c.Service.Admin) == 0 {
		return ERR_CONF_SERVICE_ADMIN
	}
	if c.Slack.Token == "" {
		return ERR_CONF_SLACK_TOKEN
	}
	if c.Slack.Syscmd == "" {
		return ERR_CONF_SLACK_SYSCMD
	}
	if c.Modules.Dir == "" {
		return ERR_CONF_MODULE_DIR
	}
	return nil
}

func isOneOf(k string, s ...string) bool {
	for i := 0; i < len(s); i++ {
		if k == s[i] {
			return true
		}
	}
	return false
}
