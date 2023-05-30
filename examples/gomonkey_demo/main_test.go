package gomonkey_demo

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"
)

func TestAdd(t *testing.T) {
	mock := func(t *testing.T) *gomonkey.Patches {
		patches := gomonkey.NewPatches()
		patches.ApplyFunc(Add, func(a, b int) int {
			return a * b
		})
		return patches
	}
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{a: 1, b: 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mk := mock(t)
			defer mk.Reset()
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDouble(t *testing.T) {
	mock := func(t *testing.T) *gomonkey.Patches {
		patches := gomonkey.NewPatches()
		patches.ApplyFunc(GetDouble, func(a int) int {
			return a
		})
		return patches
	}
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				a: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mk := mock(t)
			defer mk.Reset()
			if got := GetDouble(tt.args.a); got != tt.want {
				t.Errorf("GetDouble() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
export GOARCH=amd64 && go test -gcflags=all=-l  -v
*/
