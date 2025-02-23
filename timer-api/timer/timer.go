package timer

import (
	"os"
	"time"
)

type Response struct {
	Hostname string `json:"hostname"`
	Time     string `json:"time"`
}

func (r *Response) GetTime() Response {
	currentTime := time.Now().Format("2006-01-02T15:04:05")

	return Response{
		Hostname: os.Getenv("HOSTNAME"),
		Time:     currentTime,
	}
}
