package slackd

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Slackd struct {
	Hostname   string
	SlackToken string
	Service    service
	log        Logger
	store      Storer
	req        Requester
}

type service struct {
	Name    string
	Version string
	Contact []string
}

// New will take OPTIONAL config filename and returns a Slackd and error if any.
func New(filename string) (*Slackd, error) {
	// If filename is given, load it
	// If failed to load, an error
	if filename != "" {
		fi, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		var s Slackd
		// If cannot unmarshal the config file, return an error
		if err := json.Unmarshal(fi, &s); err != nil {
			return nil, err
		}
		// If no error, use it.
		return &s, nil
	}

	return &Slackd{
		Hostname:   ":8080",
		SlackToken: "",
		Service: service{
			Name:    "Slackd App",
			Version: "0.0.1",
			Contact: []string{},
		},
	}, nil
}

// Set will set interface items for those aren't nil.
func (s *Slackd) Set(log Logger, store Storer, req Requester) {
	if log != nil {
		s.log = log
	}
	if store != nil {
		s.store = store
	}
	if req != nil {
		s.req = req
	}
}

func (s *Slackd) Check() error {
	var err error
	err = wErrIf(s.log == nil, err, "log cannot be empty")
	err = wErrIf(s.store == nil, err, "store cannot be empty")
	err = wErrIf(s.req == nil, err, "req cannot be empty")
	return err
}

var ERR_A = errors.New("errA")
var ERR_B = errors.New("errB")
var ERR_C = errors.New("errC")


func (s *Slackd) Check2() error {
	var err error
	err = wErrs(err, ERR_A)
	err = wErrs(err, ERR_C)
	return err
}

// Run will start the service and returns an error if any.
func (s *Slackd) Run() error {
	// Check if required interface items (log, store, req) are not nil.
	if err := s.Check(); err != nil {
		return err
	}

	return nil
}
