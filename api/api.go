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

func Stringify(av **model.AsjsonValue) (string, error) {
	return stringify(av, (*av).Len)
}
