package slackd

import (
	"github.com/gonyyi/agraceful"
	"github.com/gonyyi/alog"
	"github.com/gonyyi/alog/ext"
	"github.com/gonyyi/slackd/config"
	"github.com/gonyyi/slackd/modules"
	"net/http"
	"os"
)

type Slackd struct {
	conf    config.Config
	log     alog.Logger
	modules *modules.Modules
}

// TAG for logging
var DB, HTTP, USER, SYS, REQ, RES alog.Tag

// New will take OPTIONAL config filename and returns a Slackd and error if any.
func New(filename string) (*Slackd, error) {
	if filename == "" {
		return nil, config.ERR_CONF_FILE_REQUIRED
	}

	s := new(Slackd)

	// READ CONFIGURATION
	if err := s.conf.ReadFile(filename); err != nil {
		s.log.Error()
		return s, err
	}

	// LOGGER
	s.log = alog.New(os.Stderr).SetFormatter(ext.NewFormatterTerminalColor())
	DB = s.log.NewTag("DB")
	HTTP = s.log.NewTag("HTTP")
	USER = s.log.NewTag("USER")
	SYS = s.log.NewTag("SYS")
	REQ = s.log.NewTag("REQ")
	RES = s.log.NewTag("RES")

	// CREATE MODULES
	s.modules = modules.NewModules(s.conf.Modules.Dir)

	// READY TO START
	s.log.Info(SYS).Write(INFO_SYS_READY_START)

	return s, nil
}

func (s *Slackd) Run(addr string) error {
	http.HandleFunc("/", theHandler(s))
	agraceful.IfTerm(func() {
		s.log.Fatal(SYS).
			Str("signal", agraceful.GetSignal().String()).
			Write("terminated")
		s.log.Close()
	})

	s.log.Info(HTTP|SYS).Str("addr", addr).Write(INFOS_SYS_STARTING)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		s.log.Fatal(HTTP).Err("err", err).Write(FATALS_SYS_FAILED_START)
	}

	return err
}
