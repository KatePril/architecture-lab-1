package handlers

import (
	"net/http"
)

type TimeResponse struct {
	Time string `json:"time"`
}

type Handler struct {
	Methods    []string
	Controller func(http.ResponseWriter, *http.Request)
}
