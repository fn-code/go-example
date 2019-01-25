package main

import (
	"testing"
)

func TestCheckHeader(t *testing.T) {
	var tb = []struct {
		name   string
		val    []byte
		expect error
	}{
		{"Empty byte", []byte{}, errEmptyByte},
		{"Correct header", []byte{0x81, 0x80 | byte(32), 0x7d}, nil},
		{"Invalid header mask", []byte{0x81, 0x8, 0x7d}, errMaskInvalid},
		{"invalid header mode", []byte{0x80, 0x80 | byte(32), 0x7d}, errModeInvalid},
	}

	for _, ts := range tb {
		t.Run(ts.name, func(t *testing.T) {
			err := checkHeader(ts.val)
			if err != ts.expect {
				t.Errorf("error got %v, expect %v", err, ts.expect)
			}
		})
	}

}
