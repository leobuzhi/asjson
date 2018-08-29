package parser

import (
	"strconv"

	"github.com/leobuzhi/asjson/common"
	"github.com/leobuzhi/asjson/model"
)

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
		return model.ParseRootNotSingular
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
		return model.ParseInvalidValue
	}
	ac.JSON = ac.JSON[4:]
	av.Typ = model.AsjsonNULL
	return model.ParseOK
}

func parseTrue(ac *model.AsjsonContext, av *model.AsjsonValue) error {
	if len(ac.JSON) < 4 || ac.JSON[0:4] != "true" {
		return model.ParseInvalidValue
	}
	ac.JSON = ac.JSON[4:]
	av.Typ = model.AsjsonTrue
	return model.ParseOK
}

func parseFalse(ac *model.AsjsonContext, av *model.AsjsonValue) error {
	if len(ac.JSON) < 5 || ac.JSON[0:5] != "false" {
		return model.ParseInvalidValue
	}
	ac.JSON = ac.JSON[5:]
	av.Typ = model.AsjsonFalse
	return model.ParseOK
}

func parseNumber(ac *model.AsjsonContext, av *model.AsjsonValue) error {
	slen := len(ac.JSON)
	if ac.JSON[slen-1] == '.' {
		return model.ParseInvalidValue
	}

	var i int
	if ac.JSON[i] == '-' {
		i++
	}
	if ac.JSON[i] == '0' {
		i++
	} else {
		if i < slen && !common.Isdigit1to9(ac.JSON[i]) {
			return model.ParseInvalidValue
		}
		for i++; i < slen && common.Isdigit(ac.JSON[i]); i++ {
		}
	}

	if i < slen && ac.JSON[i] == '.' {
		i++
		if i < slen && !common.Isdigit(ac.JSON[i]) {
			return model.ParseInvalidValue
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
				return model.ParseInvalidValue
			}
			for i++; i < slen && common.Isdigit(ac.JSON[i]); i++ {
			}
		}
	}
	n, err := strconv.ParseFloat(ac.JSON[0:i], 64)
	if err != nil {
		return model.ParseInvalidValue
	}
	av.Typ = model.AsjsonNumber
	av.N = n
	ac.JSON = ac.JSON[i:]
	return model.ParseOK
}

func parseString(ac *model.AsjsonContext, av *model.AsjsonValue) error {
	slen := len(ac.JSON)

	if slen < 2 || ac.JSON[0] != '"' || ac.JSON[slen-1] != '"' {
		return model.ParseMissQuotationMark
	}
	av.S = ac.JSON[1 : slen-1]
	av.Typ = model.AsjsonString
	return model.ParseOK
}

func parseArray(ac *model.AsjsonContext, av *model.AsjsonValue) error {
	alen := len(ac.JSON)
	if alen < 2 || ac.JSON[0] != '[' {
		return model.ParseMissOpenBracket
	}

	if alen == 2 && ac.JSON[1] == ']' {
		av.Typ = model.AsjsonArray
		av.Len = 0
		ac.JSON = ac.JSON[2:]
		return nil
	}

	var curr model.AsjsonValue
	var len int
	ac.JSON = ac.JSON[1:]
	for {
		var sav model.AsjsonValue
		err := parseValue(ac, &sav)
		if err != nil {
			return err
		}
		curr.Next = &sav
		len++
		if ac.JSON[0] == ',' {
			ac.JSON = ac.JSON[1:]
		} else if ac.JSON[0] == ']' {
			ac.JSON = ac.JSON[1:]
			av.Typ = model.AsjsonArray
			av.Len = len
			av.Next = curr.Next
			return nil
		} else {
			return model.ParseMissCommaOrCloseBracket
		}
	}
}

func parseValue(ac *model.AsjsonContext, av *model.AsjsonValue) error {
	if len(ac.JSON) == 0 {
		return model.ParseExpectValue
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
	default:
		return parseNumber(ac, av)
	}
}
