package flexvol

import (
	"encoding/json"
	"fmt"
)

type Capabilities struct {
	Attach bool `json:"attach"`
}

type Reply struct {
	Status       string        `json:"status"`
	Message      string        `json:"message,omitempty"`
	Device       string        `json:"device,omitempty"`
	VolumeName   string        `json:"volumeName,omitempty"`
	Capabilities *Capabilities `json:"capabilities,omitempty"`
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
		errReply := struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{"Failure", err.Error()}
		errMessage, err := json.Marshal(errReply)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(errMessage))
	}

	fmt.Println(string(message))
}
