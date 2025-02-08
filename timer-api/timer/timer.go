package timer

import "time"

type Response struct {
	Hostname string    `json:"hostname"`
	Time     time.Time `json:"time"`
}
