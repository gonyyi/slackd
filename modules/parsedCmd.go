package modules

import (
	"fmt"
	"strings"
)

type parsedCmd struct {
	ModulePath   string
	ParsedCmd    string
	ParsedArgs   string
	availableCmd []string
	isSystem     bool
	isModule     bool
	isHelp       bool
	module       *module
}

func (c parsedCmd) Module() *module {
	return c.module
}
func (c parsedCmd) IsSystem() bool {
	return c.isSystem
}
func (c parsedCmd) IsModule() bool {
	return c.isModule
}
func (c parsedCmd) IsParsed() bool {
	if c.ParsedCmd != "" {
		return true
	}
	return false
}
func (c parsedCmd) IsHelp() bool {
	return c.isHelp
}
func (c parsedCmd) InterfaceVersion() int {
	if c.module != nil {
		return c.module.Module.InterfaceVer
	}
	return -1
}
func (c parsedCmd) AvailableCmd(delimiter string) string {
	return strings.Join(c.availableCmd, delimiter)
}
func (c parsedCmd) Error() string {
	if !c.IsParsed() {
		return "Command not recognized. Use `help` for acceptable modules."
	}
	if !c.IsModule() && len(c.availableCmd) > 0 {
		return fmt.Sprintf("Unexpected sub-parsedCmd. Available format: `%s [%s]`", c.ParsedCmd, strings.Join(c.availableCmd, "|"))
	}
	return ""
}
