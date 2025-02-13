package utils

import "testing"

func TestGetEnvOrFallback(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			"env with fallback",
			[]string{"HOSTNAME", "fallback"},
			"fallback",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvOrFallback(tt.args[0], tt.args[1]); got != tt.want {
				t.Errorf("GetEnvOrFallback() = %v, want %v", got, tt.want)
			}
		})
	}
}
