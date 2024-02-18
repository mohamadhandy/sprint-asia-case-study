package commons

import "errors"

var (
	// HTTP
	HTTP_TASK_LIST                  = "HTTP|TASK_LIST"
	FAIL_GET_TASK                   = "Fail Get Task List"
	FAIL_CREATE_TASK                = "Fail Create Task"
	FAIL_CHECK_TASK                 = "Fail CHECKLIST Task"
	FAIL_UPDATE_TASK                = "Fail UPDATE Task"
	FAIL_DELETE_TASK                = "Fail DELETE Task"
	FAIL_CREATE_SUB_TASK            = "Fail CREATE SUB Task"
	FAIL_UPDATE_SUB_TASK            = "Fail UPDATE SUB Task"
	FAIL_DELETE_SUB_TASK            = "Fail DELETE SUB Task"
	FAIL_CHECKLIST_SUB_TASK         = "Fail CHECKLIST SUB Task"
	FAIL_GET_PERCENTAGE_SUB_TASK    = "Fail GET PERCENTAGE SUB Task"
	SUCCESS_CREATE_TASK             = "Success Create Task"
	SUCCESS_GET_PERCENTAGE_SUB_TASK = "Success GET PERCENTAGE SUB Task"
	SUCCESS_CREATE_SUB_TASK         = "Success Create Sub Task"
	SUCCESS_UPDATE_SUB_TASK         = "Success Update Sub Task"
	SUCCESS_DELETE_SUB_TASK         = "Success Delete Sub Task"
	SUCCESS_DELETE_TASK             = "Success Delete Task"
	SUCCESS_UPDATE_TASK             = "Success Update Task"
	SUCCESS_CHECKLIST_TASK          = "Success CHECKLIST Task"
	SUCCESS_CHECKLIST_SUBTASK       = "Success CHECKLIST Sub Task"
	SUCCESS_GET_TASK                = "Success Get Task List"
	SUCCESS_GET_TASK_HISTORY        = "Success Get Task List HISTORY"

	// USECASE CONSUMER
	USECASE_TASK = "USECASE|TASKLIST"
	// REPOSITORY MYSQL
	REPOSITORY_MYSQL_TASK = "REPOSITORY|MYSQL|TASKLIST"

	// REPOSITORY MYSQL ERROR
	ErrQuery        = errors.New("error - execute query")
	ErrPrepareQuery = errors.New("error - preparing query statement")
	ErrRowScan      = errors.New("error - scanning rows repository")

	// COMMON ERROR
	ErrInvalidPayload = errors.New("error - invalid request payload")
	ErrInternalServer = errors.New("error - internal server error")
)
