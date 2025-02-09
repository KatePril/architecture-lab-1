package handlers

import (
	"fmt"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

type Handler struct {
	Methods    []string
	Controller func(http.ResponseWriter, *http.Request)
}

func getCurrentTime() string {
	return time.Now().Format(time.RFC3339)
}

func timeHandler(response http.ResponseWriter, request *http.Request) {
	currentTime := getCurrentTime()
	fmt.Println(currentTime)
}
