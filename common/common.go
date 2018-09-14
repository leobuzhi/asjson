/*
 * @Author: Joey.Chen
 * @Date: 2018-09-10 08:25:33
 * @Last Modified by: Joey.Chen
 * @Last Modified time: 2018-09-14 22:20:01
 */
package common

// Isdigit return true when c is digit
func Isdigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// Isdigit1to9 return true when c is digit 1 to 9
func Isdigit1to9(c byte) bool {
	return c >= '1' && c <= '9'
}
