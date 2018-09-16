/*
 * @Author: Joey.Chen
 * @Date: 2018-09-10 08:25:29
 * @Last Modified by: Joey.Chen
 * @Last Modified time: 2018-09-15 22:12:18
 */
package api

import (
	"fmt"
	"strings"

	"github.com/leobuzhi/asjson/model"
)

func stringify(av **model.AsjsonValue, len int) string {
	//note(joey.chen): we ignore strings.Builder.WriteString err because it return a nil error.
	var sb strings.Builder
	switch (*av).Typ {
	case model.AsjsonNULL:
		sb.WriteString("null")
	case model.AsjsonFalse:
		sb.WriteString("false")
	case model.AsjsonTrue:
		sb.WriteString("true")
	//refefence(joey.chen):
	//http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2006/n2005.pdf
	//https://golang.org/pkg/fmt/
	case model.AsjsonNumber:
		sb.WriteString(fmt.Sprintf("%.17g", (*av).N))
	case model.AsjsonString:
		sb.WriteString("\"")
		sb.WriteString((*av).S)
		sb.WriteString("\"")
	case model.AsjsonArray:
		sb.WriteString("[")
		curr := *av
		for i := 0; i < (*av).Len && curr != nil; i++ {
			curr = (*curr).Next
			len--
			str := stringify(&curr, curr.Len)

			if i != (*av).Len-1 {
				sb.WriteString(str + ",")
			} else {
				sb.WriteString(str)
			}
		}
		if len == 0 {
			sb.WriteString("]")
		}
		*av = curr
	case model.AsjsonObject:
		sb.WriteString("{")
		curr := *av
		for i := 0; i < (*av).Len && curr != nil; i++ {
			curr = (*curr).Next
			len--
			str := stringify(&curr, curr.Len)

			if len%2 == 1 {
				sb.WriteString(str + ":")
			} else {
				if i != (*av).Len-1 {
					sb.WriteString(str + ",")
				} else {
					sb.WriteString(str)
				}
			}

		}
		if len == 0 {
			sb.WriteString("}")
		}
		*av = curr
	}

	return sb.String()
}

func stringBeautify(av **model.AsjsonValue, len, dep int) string {
	var sb strings.Builder
	switch (*av).Typ {
	case model.AsjsonArray, model.AsjsonObject:
	default:
		for i := 0; i < dep; i++ {
			sb.WriteString("  ")
		}
	}

	switch (*av).Typ {
	case model.AsjsonNULL:
		sb.Reset()
		sb.WriteString("null")
	case model.AsjsonFalse:
		sb.Reset()
		sb.WriteString("false")
	case model.AsjsonTrue:
		sb.Reset()
		sb.WriteString("true")
	//refefence(joey.chen):
	//http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2006/n2005.pdf
	//https://golang.org/pkg/fmt/
	case model.AsjsonNumber:
		sb.Reset()
		sb.WriteString(fmt.Sprintf("%.17g", (*av).N))
	case model.AsjsonString:
		sb.Reset()
		sb.WriteString("\"")
		sb.WriteString((*av).S)
		sb.WriteString("\"")
	case model.AsjsonArray:
		sb.WriteString("[\n")
		curr := *av
		for i := 0; i < (*av).Len && curr != nil; i++ {
			curr = (*curr).Next
			len--
			str := stringBeautify(&curr, curr.Len, dep+1)

			for i := 0; i < dep; i++ {
				sb.WriteString("  ")
			}
			if i != (*av).Len-1 {
				sb.WriteString(fmt.Sprintf("%s,\n", str))
			} else {
				sb.WriteString(fmt.Sprintf("%s\n", str))
			}
		}
		if len == 0 {
			for i := 0; i < dep-1; i++ {
				sb.WriteString("  ")
			}
			sb.WriteString("]")
		}
		*av = curr
	case model.AsjsonObject:
		sb.WriteString("{\n")
		curr := *av
		for i := 0; i < (*av).Len && curr != nil; i++ {
			curr = (*curr).Next
			len--
			str := stringBeautify(&curr, curr.Len, dep+1)

			if len%2 == 1 {
				for i := 0; i < dep; i++ {
					sb.WriteString("  ")
				}
				sb.WriteString(str)
				sb.WriteString(":")
			} else {
				if i != (*av).Len-1 {
					sb.WriteString(str)
					sb.WriteString(",\n")
				} else {
					sb.WriteString(str)
					sb.WriteString("\n")
				}
			}

		}
		if len == 0 {
			for i := 0; i < dep-1; i++ {
				sb.WriteString("  ")
			}
			sb.WriteString("}")
		}
		*av = curr
	}

	return sb.String()
}

//Stringify minimize json
func Stringify(av **model.AsjsonValue) string {
	return stringify(av, (*av).Len)
}

//StringBeautify beautify json
func StringBeautify(av **model.AsjsonValue) string {
	return stringBeautify(av, (*av).Len, 1)
}
