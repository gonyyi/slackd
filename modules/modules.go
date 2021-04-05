package modules

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

const (
	RESERVED_CMD_SYS     = "sys"
	RESERVED_CMD_HELP    = "help"
	RESERVED_CMD_MANNUAL = "man"

	CONFIG_FILE = "module.json"
)

// Modules hold all the module information in a directory
// Here's what Modules supposed to do
// 1. Scanning a directory and its subdirectories and find module config.
// 2. Match a parsedCmd from the slack with the module

func NewModules(baseDir string) *Modules {
	m := &Modules{
		baseDir:      baseDir,
		confFilename: CONFIG_FILE,
		modules:      make(map[string]modulesCommand),
	}
	return m
}

type modulesCommand struct {
	isModule    bool   // if the dir has config file, then it's a module, otherwise it's not.
	usage       string // if the dir has no config file, then show what's the expected by showing sub dirs
	subCommands []string
	module      *module
}

type Modules struct {
	mu           sync.RWMutex
	baseDir      string
	confFilename string
	modules      map[string]modulesCommand
}

func (m *Modules) parseCommand(cmd string) parsedCmd {
	longestMatch := ""
	longestMatchLen := 0
	longestMatchIsModule := false
	var longestMatchSubCmd []string

	for k, v := range m.modules {
		// Adding cmd and k with blank space at the end is to avoid cases such that
		// concatenated cmd and args ("cmd1 cmd2args") to be recognized incorrectly.
		if strings.HasPrefix(cmd+" ", k+" ") && len(k) > longestMatchLen {
			longestMatchLen = len(k)
			longestMatch = k
			longestMatchIsModule = v.isModule
			if v.isModule == false {
				longestMatchSubCmd = v.subCommands
			} else {
				longestMatchSubCmd = longestMatchSubCmd[:0]
			}
		}
	}

	var modulePtr *module
	if p, ok := m.modules[longestMatch]; ok {
		modulePtr = p.module
	}

	return parsedCmd{
		ModulePath:   strings.ReplaceAll(longestMatch, " ", "/"),
		ParsedCmd:    longestMatch,
		ParsedArgs:   strings.TrimPrefix(cmd, longestMatch+" "),
		availableCmd: longestMatchSubCmd,
		isModule:     longestMatchIsModule,
		module:       modulePtr,
	}
}

// ParseCommand will take user's input text, if it finds the module,
// it will return the path to the module, otherwise, it will return info.
// Return value ok will be true if it found the module, otherwise false.
func (m *Modules) ParseCommand(cmd string) parsedCmd {
	m.mu.RLock()
	defer m.mu.RUnlock()
	// Remove multiple spaces in string
	// eg. "i am     gon" --> "i am gon"
	{
		var out []string
		for _, v := range strings.Split(cmd, " ") {
			if v != "" {
				out = append(out, v)
			}
		}
		cmd = strings.Join(out, " ")
	}

	// Handles reserved command;
	var outCmd parsedCmd
	{
		tmp := strings.Split(cmd, " ")
		switch tmp[0] {
		case RESERVED_CMD_SYS:
			if len(tmp) > 1 {
				outCmd.isSystem = true
				outCmd.ParsedCmd = strings.Join(tmp[1:], " ")
				return outCmd
			}
		case RESERVED_CMD_MANNUAL, RESERVED_CMD_HELP:
			if len(tmp) > 1 {
				outCmd = m.parseCommand(strings.Join(tmp[1:], " "))
				outCmd.isHelp = true
				return outCmd
			}
		}
	}
	return m.parseCommand(cmd)
}

func (m *Modules) Scan() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Clear modules for rescanning
	for k := range m.modules {
		delete(m.modules, k)
	}

	err := filepath.Walk(m.baseDir, func(file string, info os.FileInfo, err error) error {
		baseDir := strings.TrimPrefix(path.Dir(file), m.baseDir)
		absoluteDir := path.Dir(file)
		baseFile := path.Base(file)
		cmd := strings.TrimSpace(strings.ReplaceAll(baseDir, "/", " "))

		// Checking conf filename ensures, only those who has actual modules are added
		// except for reserved keywords: system, man, help
		if tmp := strings.Split(cmd, " ")[0]; tmp == RESERVED_CMD_SYS || tmp == RESERVED_CMD_MANNUAL || tmp == RESERVED_CMD_HELP {
			// later add some error messagee
		} else if !info.IsDir() && baseFile == m.confFilename {
			tmpModule, _ := newModule(absoluteDir)
			m.modules[cmd] = modulesCommand{
				isModule: true,
				module:   tmpModule,
			}

			// If it has multi level directory, then create modules for all path
			dir3 := strings.Split(cmd, " ")
			// println("cmd:", cmd, len(dir3))
			if len(dir3) > 1 {
				for i := 0; i < len(dir3)-1; i++ {
					subCommand := strings.Join(dir3[0:i+1], " ")

					// If subCommand exist, add
					if _, ok := m.modules[subCommand]; ok {
						tmp := m.modules[subCommand]
						tmp.subCommands = append(m.modules[subCommand].subCommands, dir3[i+1])
						m.modules[subCommand] = tmp
					} else {
						// if subCommand not exist, create new
						m.modules[subCommand] = modulesCommand{
							isModule:    false,
							subCommands: []string{dir3[i+1]},
						}
					}
				}
			}
		}
		return nil
	})

	return err
}
