package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Isdigit(t *testing.T) {
	tcs := []struct {
		c byte
		b bool
	}{
		{
			'0',
			true,
		},
		{
			'1',
			true,
		},
		{
			'9',
			true,
		},
		{
			'a',
			false,
		},
	}

	for _, tc := range tcs {
		ret := Isdigit(tc.c)
		assert.Equal(t, tc.b, ret)
	}
}

func Test_Isdigit1to9(t *testing.T) {
	tcs := []struct {
		c byte
		b bool
	}{
		{
			'0',
			false,
		},
		{
			'1',
			true,
		},
		{
			'9',
			true,
		},
		{
			'a',
			false,
		},
	}

	for _, tc := range tcs {
		ret := Isdigit1to9(tc.c)
		assert.Equal(t, tc.b, ret)
	}
}
