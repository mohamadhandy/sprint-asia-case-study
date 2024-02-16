package task

import (
	"encoding/json"
	"net/http"
	"service-task-list/commons"
	http_resp "service-task-list/internal/controller/response"
	"service-task-list/internal/entity"
	"service-task-list/pkg/logger"
)

func (c *TaskRoutes) GetTaskList(w http.ResponseWriter, r *http.Request) {

	res, err := c.tu.GetTaskList(0)
	if err != nil {
		defer c.l.CreateLog(&logger.Log{
			Event:      commons.HTTP_TASK_LIST + "|GETTASKLIST",
			Method:     "GET",
			StatusCode: http.StatusInternalServerError,
			Request:    "GET TASK LIST",
			Response:   err,
			Message:    commons.FAIL_GET_TASK,
		}, logger.LVL_ERROR)
		http_resp.HttpErrorResponse(w, false, http.StatusInternalServerError, "500", err.Error())
		return
	}

	c.l.CreateLog(&logger.Log{
		Event:      commons.HTTP_TASK_LIST + "|GETTASKLIST",
		Method:     "GET",
		StatusCode: http.StatusOK,
		Response:   "Success",
		Message:    commons.SUCCESS_GET_TASK,
	}, logger.LVL_INFO)

	http_resp.HttpSuccessResponse(w, true, http.StatusOK, "200", commons.SUCCESS_GET_TASK, res)
}

func (c *TaskRoutes) GetTaskListHistory(w http.ResponseWriter, r *http.Request) {
	res, err := c.tu.GetTaskList(1)
	if err != nil {
		defer c.l.CreateLog(&logger.Log{
			Event:      commons.HTTP_TASK_LIST + "|GETTASKLIST|HISTORY",
			Method:     "GET",
			StatusCode: http.StatusInternalServerError,
			Request:    "GET TASK LIST HISTORY",
			Response:   err,
			Message:    commons.FAIL_GET_TASK,
		}, logger.LVL_ERROR)
		http_resp.HttpErrorResponse(w, false, http.StatusInternalServerError, "500", err.Error())
		return
	}

	c.l.CreateLog(&logger.Log{
		Event:      commons.HTTP_TASK_LIST + "|GETTASKLIST",
		Method:     "GET",
		StatusCode: http.StatusOK,
		Response:   "Success GET HISTORY TASK LIST",
		Message:    commons.SUCCESS_GET_TASK,
	}, logger.LVL_INFO)

	http_resp.HttpSuccessResponse(w, true, http.StatusOK, "200", commons.SUCCESS_GET_TASK_HISTORY, res)
}

func (c *TaskRoutes) CreateTask(w http.ResponseWriter, r *http.Request) {
	var payload entity.TaskRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http_resp.HttpErrorResponse(w, false, http.StatusBadRequest, "400", err.Error())
		return
	}

	err := c.tu.CreateTask(&payload)
	if err != nil {
		defer c.l.CreateLog(&logger.Log{
			Event:      commons.FAIL_CREATE_TASK + "|CREATE",
			Method:     "POST",
			StatusCode: http.StatusInternalServerError,
			Request:    payload,
			Response:   err,
			Message:    commons.FAIL_CREATE_TASK,
		}, logger.LVL_ERROR)
		http_resp.HttpErrorResponse(w, false, http.StatusInternalServerError, "500", err.Error())
		return
	}

	c.l.CreateLog(&logger.Log{
		Event:      commons.SUCCESS_CREATE_TASK + "|CREATE",
		Method:     "POST",
		StatusCode: http.StatusCreated,
		Request:    payload,
		Response:   err,
		Message:    commons.SUCCESS_CREATE_TASK,
	}, logger.LVL_INFO)

	http_resp.HttpSuccessResponse(w, true, http.StatusCreated, "201", commons.SUCCESS_CREATE_TASK, err)
}

func (c *TaskRoutes) CheckTask(w http.ResponseWriter, r *http.Request) {
	var payload entity.Task

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http_resp.HttpErrorResponse(w, false, http.StatusBadRequest, "400", err.Error())
		return
	}

	err := c.tu.CheckTask(&payload)
	if err != nil {
		defer c.l.CreateLog(&logger.Log{
			Event:      commons.FAIL_CHECK_TASK + "|CHECKLIST",
			Method:     "POST",
			StatusCode: http.StatusInternalServerError,
			Request:    payload,
			Response:   err,
			Message:    commons.FAIL_CHECK_TASK,
		}, logger.LVL_ERROR)
		http_resp.HttpErrorResponse(w, false, http.StatusInternalServerError, "500", err.Error())
		return
	}

	c.l.CreateLog(&logger.Log{
		Event:      commons.SUCCESS_CHECKLIST_TASK + "|CHECKLIST",
		Method:     http.MethodPut,
		StatusCode: http.StatusOK,
		Request:    payload,
		Response:   err,
		Message:    commons.SUCCESS_CHECKLIST_TASK,
	}, logger.LVL_INFO)

	http_resp.HttpSuccessResponse(w, true, http.StatusOK, "200", commons.SUCCESS_CHECKLIST_TASK, err)
}

func (c *TaskRoutes) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var payload entity.Task

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http_resp.HttpErrorResponse(w, false, http.StatusBadRequest, "400", err.Error())
		return
	}

	err := c.tu.UpdateTask(&payload)
	if err != nil {
		defer c.l.CreateLog(&logger.Log{
			Event:      commons.FAIL_UPDATE_TASK + "|CHECKLIST",
			Method:     "POST",
			StatusCode: http.StatusInternalServerError,
			Request:    payload,
			Response:   err,
			Message:    commons.FAIL_UPDATE_TASK,
		}, logger.LVL_ERROR)
		http_resp.HttpErrorResponse(w, false, http.StatusInternalServerError, "500", err.Error())
		return
	}

	c.l.CreateLog(&logger.Log{
		Event:      commons.SUCCESS_UPDATE_TASK + "|CHECKLIST",
		Method:     http.MethodPut,
		StatusCode: http.StatusOK,
		Request:    payload,
		Response:   err,
		Message:    commons.SUCCESS_UPDATE_TASK,
	}, logger.LVL_INFO)

	http_resp.HttpSuccessResponse(w, true, http.StatusOK, "200", commons.SUCCESS_UPDATE_TASK, err)
}
