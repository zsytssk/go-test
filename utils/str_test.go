package utils

import (
	"fmt"
	"testing"
)

func TestCamelToSnake(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"normal", args{"UserName"}, "user_name"},
		{"singleWord", args{"Username"}, "username"},
		{"withID", args{"UserID"}, "user_id"},
		{"httpPrefix", args{"HTTPServer"}, "http_server"},
		{"empty", args{""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CamelToSnake(tt.args.s)
			fmt.Printf("【%s】CamelToSnake(%q) = %q, want %q\n", tt.name, tt.args.s, got, tt.want)

			if got != tt.want {
				t.Errorf("CamelToSnake(%q) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
