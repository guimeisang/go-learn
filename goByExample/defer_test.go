package main

import (
	"os"
	"reflect"
	"testing"
)

func Test_createFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *os.File
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createFile(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
