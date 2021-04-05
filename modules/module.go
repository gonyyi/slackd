package modules

import (
	"encoding/json"
	"io/ioutil"
	"path"
)



func newModule(dir string) (*module, error) {
	m := module{}
	m.dir = dir
	err := m.Load(path.Join(dir, CONFIG_FILE))
	return &m, err
}

type module struct {
	Module struct {
		v2InterfaceVersion int    `json:"interface_version,omitempty"`
		v2EntryPoint       string `json:"entry_point,omitempty"`
		InterfaceVer       int    `json:"interface,omitempty"`
		EntryPoint         string `json:"entryPoint,omitempty"`
		Restriction        struct {
			Enable bool              `json:"enable,omitempty"`
			Users  map[string]string `json:"users,omitempty"`
		} `json:"restriction"`
	} `json:"module"`
	Info struct {
		v2LastModified string   `json:"last-modified,omitempty"`
		v2Contacts     []string `json:"contacts,omitempty"`
		Name           string   `json:"name,omitempty"`
		Version        string   `json:"version,omitempty"`
		Created        string   `json:"created,omitempty"`
		LastModified   string   `json:"lastModified,omitempty"`
		ContactEmails  []string `json:"contactEmails,omitempty"`
	} `json:"info"`
	Help struct {
		Intro         string   `json:"intro,omitempty"`
		Website       string   `json:"website,omitempty"`
		Usage         []string `json:"usage,omitempty"`
		AvgRuntimeSec int      `json:"avgRunTimeSec,omitempty"`
		MaxRuntimeSec int      `json:"maxRunTimeSec,omitempty"`
	} `json:"help"`
	dir string // this is to be filled when load
}

func (m *module) Reset() {
	m.Module.v2EntryPoint = ""
	m.Module.v2EntryPoint = ""
	m.Module.InterfaceVer = 0
	m.Module.EntryPoint = ""
	m.Module.Restriction.Enable = false
	for v := range m.Module.Restriction.Users {
		delete(m.Module.Restriction.Users, v)
	}
	m.Info.v2LastModified = ""
	m.Info.v2Contacts = m.Info.v2Contacts[:0]
	m.Info.Name = ""
	m.Info.Version = ""
	m.Info.Created = ""
	m.Info.LastModified = ""
	m.Info.ContactEmails = m.Info.ContactEmails[:0]
	m.Help.Intro = ""
	m.Help.Website = ""
	m.Help.Usage = m.Help.Usage[:0]
	m.Help.AvgRuntimeSec = 0
	m.Help.MaxRuntimeSec = 0
}

func (m *module) Load(filename string) error {
	m.Reset()
	fi, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(fi, m); err != nil {
		return err
	}

	// For backward compatibility, if newer field of version doesn't populates,
	// copy everything from older version into new.
	// note that, custom going to and from the module should be compatible as well.
	if m.Module.InterfaceVer == 0 {
		m.Module.InterfaceVer = m.Module.v2InterfaceVersion
		m.Module.EntryPoint = m.Module.v2EntryPoint
		m.Info.LastModified = m.Info.v2LastModified
		m.Info.ContactEmails = m.Info.v2Contacts
	}

	return nil
}

func (m *module) Save(filename string) error {
	b, err := json.MarshalIndent(m, "", "   ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, b, 0755)
}
