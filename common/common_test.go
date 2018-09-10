/*
 * @Author: Joey.Chen
 * @Date: 2018-09-10 08:25:31
 * @Last Modified by: Joey.Chen
 * @Last Modified time: 2018-09-10 08:26:22
 */
package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
