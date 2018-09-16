/*
 * @Author: Joey.Chen
 * @Date: 2018-09-10 08:25:48
 * @Last Modified by: Joey.Chen
 * @Last Modified time: 2018-09-15 10:55:26
 */
package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
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

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		stdin, err := os.OpenFile("./testdata/testdata1.json", os.O_RDONLY, 0664)
		os.Stdin = stdin
		assert.Equal(b, nil, err)
		main()
	}
}
