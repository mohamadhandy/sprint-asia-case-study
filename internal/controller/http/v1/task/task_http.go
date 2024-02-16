package task

import (
	"encoding/json"
	"net/http"
	"service-task-list/commons"
	http_resp "service-task-list/internal/controller/response"
	"service-task-list/internal/entity"
	"service-task-list/pkg/logger"
	"strconv"

	"github.com/gorilla/mux"
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
	param := mux.Vars(r)
	taskId := param["task_id"]
	intTaskId, err := strconv.Atoi(taskId)
	if err != nil {
		defer c.l.CreateLog(&logger.Log{
			Event:      commons.FAIL_CHECK_TASK + "|CHECKLIST|CONVERT_ID",
			Method:     http.MethodPut,
			StatusCode: http.StatusInternalServerError,
			Request:    intTaskId,
			Response:   err,
			Message:    commons.FAIL_CHECK_TASK,
		}, logger.LVL_ERROR)
		http_resp.HttpErrorResponse(w, false, http.StatusInternalServerError, "500", err.Error())
		return
	}

	err = c.tu.CheckTask(intTaskId)
	if err != nil {
		defer c.l.CreateLog(&logger.Log{
			Event:      commons.FAIL_CHECK_TASK + "|CHECKLIST",
			Method:     http.MethodPut,
			StatusCode: http.StatusInternalServerError,
			Request:    intTaskId,
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
		Request:    intTaskId,
		Response:   err,
		Message:    commons.SUCCESS_CHECKLIST_TASK,
	}, logger.LVL_INFO)

	http_resp.HttpSuccessResponse(w, true, http.StatusOK, "200", commons.SUCCESS_CHECKLIST_TASK, err)
}

func (c *TaskRoutes) UpdateTask(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	taskId := param["task_id"]
	intTaskId, err := strconv.Atoi(taskId)
	if err != nil {
		defer c.l.CreateLog(&logger.Log{
			Event:      commons.FAIL_UPDATE_TASK + "|CONVERT_ID",
			Method:     http.MethodPut,
			StatusCode: http.StatusInternalServerError,
			Request:    intTaskId,
			Response:   err,
			Message:    commons.FAIL_UPDATE_TASK,
		}, logger.LVL_ERROR)
		http_resp.HttpErrorResponse(w, false, http.StatusInternalServerError, "500", err.Error())
		return
	}
	var payload entity.Task

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http_resp.HttpErrorResponse(w, false, http.StatusBadRequest, "400", err.Error())
		return
	}

	if intTaskId == payload.ID {
		err := c.tu.UpdateTask(&payload)
		if err != nil {
			defer c.l.CreateLog(&logger.Log{
				Event:      commons.FAIL_UPDATE_TASK,
				Method:     http.MethodPut,
				StatusCode: http.StatusInternalServerError,
				Request:    payload,
				Response:   err,
				Message:    commons.FAIL_UPDATE_TASK,
			}, logger.LVL_ERROR)
			http_resp.HttpErrorResponse(w, false, http.StatusInternalServerError, "500", err.Error())
			return
		}

		c.l.CreateLog(&logger.Log{
			Event:      commons.SUCCESS_UPDATE_TASK,
			Method:     http.MethodPut,
			StatusCode: http.StatusOK,
			Request:    payload,
			Response:   err,
			Message:    commons.SUCCESS_UPDATE_TASK,
		}, logger.LVL_INFO)
		http_resp.HttpSuccessResponse(w, true, http.StatusOK, "200", commons.SUCCESS_UPDATE_TASK, err)
		return
	} else {
		defer c.l.CreateLog(&logger.Log{
			Event:      commons.FAIL_UPDATE_TASK,
			Method:     http.MethodPut,
			StatusCode: http.StatusBadRequest,
			Request:    payload,
			Response:   err,
			Message:    commons.FAIL_UPDATE_TASK,
		}, logger.LVL_ERROR)
		http_resp.HttpErrorResponse(w, false, http.StatusBadRequest, "400", "ID not match")
		return
	}

}

func (c *TaskRoutes) DeleteTask(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	taskId := param["task_id"]
	intTaskId, err := strconv.Atoi(taskId)
	if err != nil {
		defer c.l.CreateLog(&logger.Log{
			Event:      commons.FAIL_DELETE_TASK + "|CONVERT_ID",
			Method:     "POST",
			StatusCode: http.StatusInternalServerError,
			Request:    intTaskId,
			Response:   err,
			Message:    commons.FAIL_DELETE_TASK,
		}, logger.LVL_ERROR)
		http_resp.HttpErrorResponse(w, false, http.StatusInternalServerError, "500", err.Error())
		return
	}

	err = c.tu.DeleteTask(intTaskId)
	if err != nil {
		defer c.l.CreateLog(&logger.Log{
			Event:      commons.FAIL_DELETE_TASK + "|DELETE",
			Method:     "POST",
			StatusCode: http.StatusInternalServerError,
			Request:    taskId,
			Response:   err,
			Message:    commons.FAIL_DELETE_TASK,
		}, logger.LVL_ERROR)
		http_resp.HttpErrorResponse(w, false, http.StatusInternalServerError, "500", err.Error())
		return
	}

	c.l.CreateLog(&logger.Log{
		Event:      commons.SUCCESS_DELETE_TASK + "|DELETE",
		Method:     http.MethodDelete,
		StatusCode: http.StatusOK,
		Request:    taskId,
		Response:   err,
		Message:    commons.SUCCESS_DELETE_TASK,
	}, logger.LVL_INFO)

	http_resp.HttpSuccessResponse(w, true, http.StatusOK, "200", commons.SUCCESS_DELETE_TASK, err)
}
