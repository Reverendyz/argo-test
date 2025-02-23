package timer

import (
	"os"
	"testing"
	"time"
)

func TestResponse_GetTime(t *testing.T) {
	mockHostname := "PTWGVP1RNAKTVYT2"
	os.Setenv("HOSTNAME", mockHostname)
	defer os.Unsetenv("HOSTNAME")

	tests := []struct {
		name   string
		fields Response
		want   Response
	}{
		{
			"timer with no hostname",
			Response{
				Hostname: " ",
				Time:     time.Now().Format("2006-01-02T15:04:05"),
			},
			Response{
				Hostname: mockHostname,
				Time:     time.Now().Format("2006-01-02T15:04:05"),
			},
		},
		{
			"timer with hostname",
			Response{
				Hostname: "SomeGarbage",
				Time:     time.Now().Format("2006-01-02T15:04:05"),
			},
			Response{
				Hostname: mockHostname,
				Time:     time.Now().Format("2006-01-02T15:04:05"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				Hostname: tt.fields.Hostname,
				Time:     tt.fields.Time,
			}

			got := r.GetTime()

			if got.Hostname != tt.want.Hostname {
				t.Errorf("Response.GetTime() Hostname = %v, want %v", got.Hostname, tt.want.Hostname)
			}

			if got.Time != tt.want.Time {
				t.Errorf("Response.GetTime() Time = %v, want %v", got.Time, tt.want.Time)
			}
		})
	}
}
