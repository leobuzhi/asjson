/*
 * @Author: Joey.Chen
 * @Date: 2018-09-10 08:25:48
 * @Last Modified by: Joey.Chen
 * @Last Modified time: 2018-09-15 00:24:31
 */
package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
		min  bool
	}{
		{
			"with min flag",
			true,
		},
		{
			"without min flag",
			false,
		},
	}
	for _, tt := range tests {
		*min = tt.min
		stdin, err := os.OpenFile("./testdata/testdata1.json", os.O_RDONLY, 0664)
		assert.Equal(t, nil, err)
		os.Stdin = stdin
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
