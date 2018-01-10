package main

import "testing"

func Test_formatURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatURL(tt.args.url); got != tt.want {
				t.Errorf("formatURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
