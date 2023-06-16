package util

import (
	"fmt"
	"testing"
)

func TestGetRandString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantLen int
	}{
		{name: "case1", args: args{length: 6}, wantLen: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetRandString(tt.args.length)
			if len(got) != tt.wantLen {
				t.Errorf("GetRandString() = %v, want %v", got, tt.want)
			}

			fmt.Println(got)
		})
	}
}

func TestGetRandStringWithCharset(t *testing.T) {
	type args struct {
		length  int
		charset string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantLen int
	}{
		{name: "case1", args: args{length: 6, charset: "0x0x0x"}, wantLen: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetRandStringWithCharset(tt.args.length, tt.args.charset)
			if len(got) != tt.wantLen {
				t.Errorf("GetRandStringWithCharset() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}
