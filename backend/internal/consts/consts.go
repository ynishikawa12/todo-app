package consts

const (
	UserURL          = "/users"
	TaskURL          = "/tasks"
	TaskTagURL       = "/task-tags"
	UserIDParam      = "user-id"
	TaskIDParam      = "task-id"
	TaskTagIDParam   = "task-tag-id"
	UserURLWithID    = UserURL + "/:" + UserIDParam
	TaskURLWithID    = TaskURL + "/:" + TaskIDParam
	TaskTagURLWithID = TaskTagURL + "/:" + TaskTagIDParam
)
