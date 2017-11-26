package template

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Status        string      `json:"status"`
	MessageStatus []string    `json:"message_status,omitempty"`
	MessageError  []string    `json:"message_error,omitempty"`
	Config        interface{} `json:"config"`
	Data          interface{} `json:"data"`
	ProcessTime   string      `json:"server_process_time"`
}

func ResponseError(w http.ResponseWriter, status int, messageError []string, processTime time.Time) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	r := Response{
		Status:       http.StatusText(status),
		MessageError: messageError,
		ProcessTime:  fmt.Sprintf("%f", time.Since(processTime).Seconds()),
	}

	//pretty print
	// not recommeded, extra bytes for space LOL
	responseRaw, _ := json.MarshalIndent(r, "", "  ")

	//no space print
	// responseRaw, _ := json.Marshal(r)
	w.Write(responseRaw)
}

func ResponseOK(w http.ResponseWriter, data interface{}, messageStatus []string, processTime time.Time) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	status := http.StatusOK
	w.WriteHeader(status)

	r := Response{
		Status:        http.StatusText(status),
		Data:          data,
		MessageStatus: messageStatus,
		ProcessTime:   fmt.Sprintf("%f", time.Since(processTime).Seconds()),
	}
	responseRaw, _ := json.MarshalIndent(r, "", "  ")
	w.Write(responseRaw)
}
