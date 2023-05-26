package service

import (
	"testing"
)

func TestPrintFixedColumnVersion(t *testing.T) {
	type args struct {
		data []Version
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{data: []Version{
			{"v1.0.0"}, {"v1.2.3"}, {"v1.5.2"}, {"v2.0.0"}, {"v2.1.1"},
			{"v2.3.0"}, {"v3.0.0"}, {"v3.5.1"}, {"v4.2.0"}, {"v4.6.3"},
			{"v5.0.0"}, {"v5.1.2"}, {"v6.0.0"}, {"v6.3.2"}, {"v7.1.0"},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintFixedColumnVersion(tt.args.data)
		})
	}
}

func TestPrintMagenta(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "china", args: args{text: "Hello, China!"}},
		{name: "world", args: args{text: "Hello, World!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintMagenta(tt.args.text)
		})
	}
}
