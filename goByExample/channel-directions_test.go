package main

import "testing"

func Test_ping(t *testing.T) {
	type args struct {
		pings chan<- string
		msg   string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ping(tt.args.pings, tt.args.msg)
		})
	}
}

func Test_pong(t *testing.T) {
	type args struct {
		pings <-chan string
		pongs chan<- string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pong(tt.args.pings, tt.args.pongs)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
