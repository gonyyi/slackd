package slackd

const (
	INFO_SYS_READY_START          = "system ready to start"
	INFO_RECEIVED_SLACK_CHALLENGE = "received a challenge request"

	INFOS_SYS_STARTING = "system starting"

	WARNS_UNEXPECTED_EXCEPTION = "unexpected exception"

	ERRS_DB_START_FAILED          = "failed to start DB"
	ERRS_FAILED_READ_REQ_BODY     = "failed to read request body"
	ERRS_FAILED_LOAD_INCOMING_MSG = "failed to load incoming message"

	FATALS_SYS_FAILED_START = "system failed to start"
)
