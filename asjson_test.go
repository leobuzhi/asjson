/*
 * @Author: Joey.Chen
 * @Date: 2018-09-10 08:25:48
 * @Last Modified by: Joey.Chen
 * @Last Modified time: 2018-09-11 22:48:38
 */
package main

import "testing"

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
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
