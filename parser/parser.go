/*
 * @Author: Joey.Chen
 * @Date: 2018-09-10 08:25:11
 * @Last Modified by: Joey.Chen
 * @Last Modified time: 2018-09-14 22:32:13
 */
package parser

import (
	"strconv"

	"github.com/leobuzhi/asjson/common"
	"github.com/leobuzhi/asjson/model"
)

//Parse parse json to AsjsonValue return parse error
func Parse(json string, av *model.AsjsonValue) error {
	var ac model.AsjsonContext
	ac.JSON = json
	parseWhitespace(&ac)
	err := parseValue(&ac, av)
	if err != nil {
		return err
	}
	parseWhitespace(&ac)
	if len(ac.JSON) > 0 {
		av.Typ = model.AsjsonNAT
		return model.ErrParseRootNotSingular
	}
	return err
}

func parseWhitespace(ac *model.AsjsonContext) {
	var j int
	for i := 0; i < len(ac.JSON); i++ {
		if ac.JSON[i] == ' ' || ac.JSON[i] == '\t' || ac.JSON[i] == '\n' || ac.JSON[i] == '\r' {
			j++
			continue
		}
		break
	}
	ac.JSON = ac.JSON[j:]
}

func parseNull(ac *model.AsjsonContext, av *model.AsjsonValue) error {
	if len(ac.JSON) < 4 || ac.JSON[0:4] != "null" {
		return model.ErrParseInvalidValue
	}
	ac.JSON = ac.JSON[4:]
	av.Typ = model.AsjsonNULL
	return model.ParseOK
}

func parseTrue(ac *model.AsjsonContext, av *model.AsjsonValue) error {
	if len(ac.JSON) < 4 || ac.JSON[0:4] != "true" {
		return model.ErrParseInvalidValue
	}
	ac.JSON = ac.JSON[4:]
	av.Typ = model.AsjsonTrue
	return model.ParseOK
}

func parseFalse(ac *model.AsjsonContext, av *model.AsjsonValue) error {
	if len(ac.JSON) < 5 || ac.JSON[0:5] != "false" {
		return model.ErrParseInvalidValue
	}
	ac.JSON = ac.JSON[5:]
	av.Typ = model.AsjsonFalse
	return model.ParseOK
}

func parseNumber(ac *model.AsjsonContext, av *model.AsjsonValue) error {
	slen := len(ac.JSON)
	if ac.JSON[slen-1] == '.' {
		return model.ErrParseInvalidValue
	}

	var i int
	if ac.JSON[i] == '-' {
		i++
	}
	if ac.JSON[i] == '0' {
		i++
	} else {
		if i < slen && !common.Isdigit1to9(ac.JSON[i]) {
			return model.ErrParseInvalidValue
		}
		for i++; i < slen && common.Isdigit(ac.JSON[i]); i++ {
		}
	}

	if i < slen && ac.JSON[i] == '.' {
		i++
		if i < slen && !common.Isdigit(ac.JSON[i]) {
			return model.ErrParseInvalidValue
		}
		for i++; i < slen && common.Isdigit(ac.JSON[i]); i++ {
		}
	}
	if i < slen {
		if ac.JSON[i] == 'e' || ac.JSON[i] == 'E' {
			i++
			if i < slen {
				if ac.JSON[i] == '+' || ac.JSON[i] == '-' {
					i++
				}
			}
			if i < slen && !common.Isdigit(ac.JSON[i]) {
				return model.ErrParseInvalidValue
			}
			for i++; i < slen && common.Isdigit(ac.JSON[i]); i++ {
			}
		}
	}
	n, err := strconv.ParseFloat(ac.JSON[0:i], 64)
	if err != nil {
		return model.ErrParseInvalidValue
	}
	av.Typ = model.AsjsonNumber
	av.N = n
	ac.JSON = ac.JSON[i:]
	return model.ParseOK
}

func parseString(ac *model.AsjsonContext, av *model.AsjsonValue) error {
	slen := len(ac.JSON)

	if slen < 2 || ac.JSON[0] != '"' {
		return model.ErrParseMissQuotationMark
	}
	var i int
	for i = 1; i < len(ac.JSON); i++ {
		if ac.JSON[i] == '"' {
			break
		}
	}

	av.S = ac.JSON[1:i]
	av.Typ = model.AsjsonString
	ac.JSON = ac.JSON[i+1:]
	return model.ParseOK
}

func parseArray(ac *model.AsjsonContext, av *model.AsjsonValue) error {
	alen := len(ac.JSON)
	if alen < 2 || ac.JSON[0] != '[' {
		return model.ErrParseMissOpenBracket
	}

	curr := av
	var len int
	ac.JSON = ac.JSON[1:]
	parseWhitespace(ac)
	if ac.JSON[0] == ']' {
		ac.JSON = ac.JSON[1:]
		av.Typ = model.AsjsonArray
		av.Len = 0
		return nil
	}

	for {
		parseWhitespace(ac)
		sav := new(model.AsjsonValue)
		err := parseValue(ac, sav)
		if err != nil {
			return err
		}
		parseWhitespace(ac)

		curr.Next = sav
		for ; curr.Next != nil; curr = curr.Next {
		}

		len++

		if ac.JSON == "" {
			return model.ErrParseMissCloseBracket
		}

		if ac.JSON[0] == ',' {
			ac.JSON = ac.JSON[1:]
		} else if ac.JSON[0] == ']' {
			ac.JSON = ac.JSON[1:]
			av.Typ = model.AsjsonArray
			av.Len = len
			return nil
		} else {
			return model.ErrParseMissCommaOrCloseBracket
		}
	}
}

func parseObject(ac *model.AsjsonContext, av *model.AsjsonValue) error {
	alen := len(ac.JSON)
	if alen < 2 || ac.JSON[0] != '{' {
		return model.ErrParseMissOpenBrace
	}

	curr := av
	var len int
	ac.JSON = ac.JSON[1:]
	parseWhitespace(ac)
	if ac.JSON[0] == '}' {
		ac.JSON = ac.JSON[1:]
		av.Typ = model.AsjsonObject
		av.Len = 0
		return nil
	}

	for {
		parseWhitespace(ac)
		sav := new(model.AsjsonValue)
		err := parseString(ac, sav)
		if err != nil {
			return err
		}
		parseWhitespace(ac)
		curr.Next = sav
		curr = curr.Next

		if ac.JSON == "" || ac.JSON[0] != ':' {
			return model.ErrParseMissColon
		}
		ac.JSON = ac.JSON[1:]
		parseWhitespace(ac)
		sav = new(model.AsjsonValue)
		err = parseValue(ac, sav)
		if err != nil {
			return err
		}
		parseWhitespace(ac)

		curr.Next = sav
		for ; curr.Next != nil; curr = curr.Next {
		}
		len += 2

		if ac.JSON == "" {
			return model.ErrParseMissCloseBrace
		}

		if ac.JSON[0] == ',' {
			ac.JSON = ac.JSON[1:]
		} else if ac.JSON[0] == '}' {
			ac.JSON = ac.JSON[1:]
			av.Typ = model.AsjsonObject
			av.Len = len
			return nil
		} else {
			return model.ErrParseMissCommaOrCloseBrace
		}
	}
}

func parseValue(ac *model.AsjsonContext, av *model.AsjsonValue) error {
	if len(ac.JSON) == 0 {
		return model.ErrParseExpectValue
	}
	switch ac.JSON[0] {
	case 'n':
		return parseNull(ac, av)
	case 't':
		return parseTrue(ac, av)
	case 'f':
		return parseFalse(ac, av)
	case '"':
		return parseString(ac, av)
	case '[':
		return parseArray(ac, av)
	case '{':
		return parseObject(ac, av)
	default:
		return parseNumber(ac, av)
	}
}
