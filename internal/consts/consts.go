package consts

const (
	//ENV
	DefaultPrefixForEnv = "CUSTOMER"
	ServiceHostKey      = "HTTP_HOST"
	ServicePortKey      = "HTTP_PORT"
	ServiceDBUrlKey     = "DB_URL"
	ServiceLogOutputKey = "LOG_OUTPUT"
	ServiceLogLevelKey  = "LOG_LEVEL"
	ServiceLogFormatKey = "LOG_FORMAT"
	DefaultDateLayout   = "2006-01-02"

	//actions
	ActionCreate  = "create"
	ActionUpdate  = "update"
	ActionSearch  = "search"
	ActionDelete  = "delete"
	ActionGetByID = "get by id"
	ActionGetAll  = "get all"
)
