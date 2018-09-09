package api

import (
	"fmt"

	"github.com/leobuzhi/asjson/model"
)

func stringify(av **model.AsjsonValue, len int) (string, error) {
	var ret string
	switch (*av).Typ {
	case model.AsjsonNULL:
		ret = "null"
	case model.AsjsonFalse:
		ret = "false"
	case model.AsjsonTrue:
		ret = "true"
	//refefence(joey.chen):
	//http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2006/n2005.pdf
	//https://golang.org/pkg/fmt/
	case model.AsjsonNumber:
		ret = fmt.Sprintf("%.17g", (*av).N)
	case model.AsjsonString:
		ret = fmt.Sprintf("\"%s\"", (*av).S)
	case model.AsjsonArray:
		ret += "["
		curr := *av
		for i := 0; i < (*av).Len && curr != nil; i++ {
			curr = (*curr).Next
			len--
			str, err := stringify(&curr, curr.Len)
			if err != nil {
				return "", err
			}

			if i != (*av).Len-1 {
				ret += str + ","
			} else {
				ret += str
			}
		}
		if len == 0 {
			ret += "]"
		}
		*av = curr
	case model.AsjsonObject:
		ret += "{"
		curr := *av
		for i := 0; i < (*av).Len && curr != nil; i++ {
			curr = (*curr).Next
			len--
			str, err := stringify(&curr, curr.Len)
			if err != nil {
				return "", err
			}

			if len%2 == 1 {
				ret += str + ":"
			} else {
				if i != (*av).Len-1 {
					ret += str + ","
				} else {
					ret += str
				}
			}

		}
		if len == 0 {
			ret += "}"
		}
		*av = curr
	}

	return ret, nil
}

func stringBeautify(av **model.AsjsonValue, len, dep int) (string, error) {
	var ret string
	switch (*av).Typ {
	case model.AsjsonArray, model.AsjsonObject:
	default:
		for i := 0; i < dep; i++ {
			ret += "  "
		}
	}

	switch (*av).Typ {
	case model.AsjsonNULL:
		ret = "null"
	case model.AsjsonFalse:
		ret = "false"
	case model.AsjsonTrue:
		ret = "true"
	//refefence(joey.chen):
	//http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2006/n2005.pdf
	//https://golang.org/pkg/fmt/
	case model.AsjsonNumber:
		ret = fmt.Sprintf("%.17g", (*av).N)
	case model.AsjsonString:
		ret = fmt.Sprintf("\"%s\"", (*av).S)
	case model.AsjsonArray:
		ret += "[\n"
		curr := *av
		for i := 0; i < (*av).Len && curr != nil; i++ {
			curr = (*curr).Next
			len--
			str, err := stringBeautify(&curr, curr.Len, dep+1)
			if err != nil {
				return "", err
			}

			for i := 0; i < dep; i++ {
				ret += "  "
			}
			if i != (*av).Len-1 {
				ret += fmt.Sprintf("%s,\n", str)
			} else {
				ret += fmt.Sprintf("%s\n", str)
			}
		}
		if len == 0 {
			for i := 0; i < dep-1; i++ {
				ret += "  "
			}
			ret += "]"
		}
		*av = curr
	case model.AsjsonObject:
		ret += "{\n"
		curr := *av
		for i := 0; i < (*av).Len && curr != nil; i++ {
			curr = (*curr).Next
			len--
			str, err := stringBeautify(&curr, curr.Len, dep+1)
			if err != nil {
				return "", err
			}

			if len%2 == 1 {
				for i := 0; i < dep; i++ {
					ret += "  "
				}
				ret += fmt.Sprintf("%s:", str)
			} else {
				if i != (*av).Len-1 {
					ret += fmt.Sprintf("%s,\n", str)
				} else {
					ret += fmt.Sprintf("%s\n", str)
				}
			}

		}
		if len == 0 {
			for i := 0; i < dep-1; i++ {
				ret += "  "
			}
			ret += "}"
		}
		*av = curr
	}

	return ret, nil
}

func Stringify(av **model.AsjsonValue) (string, error) {
	return stringify(av, (*av).Len)
}

func StringBeautify(av **model.AsjsonValue) (string, error) {
	return stringBeautify(av, (*av).Len, 1)
}
