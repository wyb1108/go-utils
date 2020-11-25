package util

import "testing"

func TestGetCurrentDay(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"test",
			"20200907",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCurrentDay(); got != tt.want {
				t.Errorf("GetCurrentDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
