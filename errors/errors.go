package errors

import (
	"encoding/json"
	"net/http"
)

type CustomError struct {
	Msg  string
	Code int
}

func (e CustomError) Error() string {
	return e.Msg
}

func NewNotFound(message string) *CustomError {
	return &CustomError{message, 404}
}

func NewBadRequest(message string) *CustomError {
	return &CustomError{message, 400}
}

func NewUnauthorized(message string) *CustomError {
	return &CustomError{message, 401}
}

func NewForbidden(message string) *CustomError {
	return &CustomError{message, 403}
}

func NewInternal(message string) *CustomError {
	return &CustomError{message, 500}
}

func HandleError(w http.ResponseWriter, err *CustomError) {
	errJSON := struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)

	data, marshalErr := json.Marshal(errJSON)
	if marshalErr != nil {
		http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		return
	}

	_, writeErr := w.Write(data)
	if writeErr != nil {
		http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		return
	}

	return
}
