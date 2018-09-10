/*
 * @Author: Joey.Chen
 * @Date: 2018-09-10 08:25:33
 * @Last Modified by: Joey.Chen
 * @Last Modified time: 2018-09-10 08:26:16
 */
package common

func Isdigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func Isdigit1to9(c byte) bool {
	return c >= '1' && c <= '9'
}
