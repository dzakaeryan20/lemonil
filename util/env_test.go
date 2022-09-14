package util

import (
	"testing"

	"github.com/dzakaeryan20/lemonilo/model"
)

func TestGetEnv(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "success",
			want: model.LocalEnv,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnv(); got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
