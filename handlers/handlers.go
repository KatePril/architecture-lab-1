package handlers

import (
	"encoding/json"
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
	answer, err := json.Marshal(TimeResponse{ currentTime })
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(answer)
}

var Handlers = map[string] Handler {
	"/time" : { []string{ http.MethodGet }, timeHandler },
}
