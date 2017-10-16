package main

import (
	"reflect"
	"testing"
)

func Test_trcDecrypt(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			"yo",
			args{
				[]byte{
					0x4B, 0xD1, 0x10, 0x13, 0x51, 0x50, 0x13, 0xD1, 0x51, 0xCA, 0x95, 0xCF, 0x0E, 0xCF, 0x88, 0x12,
					0x12, 0x95, 0xCC, 0x83, 0x45, 0xC4, 0x81, 0x80, 0xC7, 0x46, 0x06, 0x95, 0x8F, 0xC5, 0x06, 0x84,
					0x86, 0xC4, 0x01, 0x95, 0xC7, 0x41, 0x95, 0x0C, 0xC7, 0x06, 0xC7, 0x41, 0x87, 0xC4, 0x84, 0x16,
					0xDE, 0x1F,
				},
			},
			[]byte("[16:37:13] INIT>> Exception Handler is Finished.\r\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trcDecrypt(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trcDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}