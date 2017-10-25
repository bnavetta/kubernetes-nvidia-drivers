package flexvol

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Reply struct {
	Status     string `json:"status"`
	Message    string `json:"message,omitempty"`
	Device     string `json:"device,omitempty"`
	VolumeName string `json:"volumeName,omitempty"`
	//Attached   bool
}

const (
	StatusSuccess      = "Success"
	StatusFailure      = "Failure"
	StatusNotSupported = "Not supported"
)

func Failure(message string) Reply {
	return Reply{
		Status:  StatusFailure,
		Message: message,
	}
}

func Success(message string) Reply {
	return Reply{
		Status:  StatusSuccess,
		Message: message,
	}
}

func Log(reply Reply) {
	message, err := json.Marshal(reply)
	if err != nil {
		fmt.Printf(`{"status": "Failure", "message": "%v"}`, strings.Replace(err.Error(), `"`, `\"`, -1))
		//os.Exit(1)
	}

	os.Stdout.Write(message)
	//if reply.Status == StatusFailure {
	//	os.Exit(1)
	//} else {
	//	os.Exit(0)
	//}
}
