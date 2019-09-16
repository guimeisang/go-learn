package main

import "testing"

func TestByLength_Len(t *testing.T) {
	tests := []struct {
		name string
		s    ByLength
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("ByLength.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByLength_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		s    ByLength
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestByLength_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		s    ByLength
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("ByLength.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}
