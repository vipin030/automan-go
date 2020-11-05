package errors

import (
	"fmt"
	"net/http"
)

// APIError interface
type APIError interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type apiErr struct {
	message string        `json:"message"`
	status  int           `json:"status_code"`
	error   string        `json:"error"`
	causes  []interface{} `json:"causes"`
}

func (e apiErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: [%v]", e.message, e.status, e.error, e.causes)
}

func (e apiErr) Message() string {
	return e.message
}

func (e apiErr) Status() int {
	return e.status
}

func (e apiErr) Causes() []interface{} {
	return e.causes
}

// CustomError for all type of error
func CustomError(message string, status int, err string, causes []interface{}) APIError {
	return apiErr{
		message: message,
		status:  status,
		error:   err,
		causes:  causes,
	}
}

func ValidationError(message string, err string) APIError {
	return apiErr{
		message: message,
		status:  http.StatusUnprocessableEntity,
		error:   err,
	}
}

// BadRequestError 400
func BadRequestError(message string) APIError {
	return apiErr{
		message: message,
		status:  http.StatusBadRequest,
		error:   "bad_request",
	}
}

// UnAutherizedError 401
func UnAutherizedError(message string) APIError {
	return apiErr{
		message: message,
		status:  http.StatusUnauthorized,
		error:   "unauthorized",
	}
}

// NotFoundError - 404
func NotFoundError(message string) APIError {
	return apiErr{
		message: message,
		status:  http.StatusNotFound,
		error:   "not_found",
	}
}

//InternalServerError 500
func InternalServerError(message string, err error) APIError {
	result := apiErr{
		message: message,
		status:  http.StatusInternalServerError,
		error:   "internal_server_error",
	}
	if err != nil {
		result.causes = append(result.causes, err.Error())
	}
	return result
}
