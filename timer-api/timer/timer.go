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
	currentDate := time.Now().Format("2006-01-02")

	return Response{
		Hostname: os.Getenv("HOSTNAME"),
		Time:     currentDate,
	}
}
