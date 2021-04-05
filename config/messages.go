package config

import (
	"github.com/gonyyi/atype"
)

// Global Constants/Variables
const (
	ERR_CONF_FILE_REQUIRED            = atype.ErrorStr("Config file is required")
	ERR_BAD_CONF_FILE                 = atype.ErrorStr("bad Config file")
	ERR_CONF_SYSTEM_DB_MISSING        = atype.ErrorStr("system db_file is missing")
	ERR_CONF_LOGGING_FILENAME_MISSING = atype.ErrorStr("filename for logging is missing")
	ERR_CONF_LOGGING_INCORRECT_LEVEL  = atype.ErrorStr("incorrect logging level value (acceptable values: TRACE, DEBUG, INFO, WARN, ERROR, FATAL)")
	ERR_CONF_LOGGING_INVALID_MAXSIZE  = atype.ErrorStr("logging max_size_kb is missing")
	ERR_CONF_SERVICE_NAME             = atype.ErrorStr("service name is missing")
	ERR_CONF_SERVICE_VERSION          = atype.ErrorStr("service version is missing")
	ERR_CONF_SERVICE_HOST             = atype.ErrorStr("service host value is missing")
	ERR_CONF_SERVICE_ADMIN            = atype.ErrorStr("service admin is missing")
	ERR_CONF_SLACK_TOKEN              = atype.ErrorStr("slack token missing")
	ERR_CONF_SLACK_SYSCMD             = atype.ErrorStr("slack syscmd missing")
	ERR_CONF_MODULE_DIR               = atype.ErrorStr("module directory name missing")
	ERR_CONF_MODULE_CONF_FILE         = atype.ErrorStr("module Config filename missing")
)
