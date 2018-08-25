package parse

import (
	"github.com/leobuzhi/asjson/model"
	"strconv"
	"strings"
)

func Parse(json string, av *model.AsjsonValue) error {
	var ac model.AsjsonContext
	ac.JSON = json
	parseWhitespace(&ac)
	err := parseValue(&ac, av)
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
		av.Typ = model.AsjsonNAT
		return model.ParseInvalidValue
	}
	ac.JSON = ac.JSON[4:]
	av.Typ = model.AsjsonNULL
	return model.ParseOK
}

func parseTrue(ac *model.AsjsonContext, av *model.AsjsonValue) error {
	if len(ac.JSON) < 4 || ac.JSON[0:4] != "true" {
		av.Typ = model.AsjsonNAT
		return model.ParseInvalidValue
	}
	ac.JSON = ac.JSON[4:]
	av.Typ = model.AsjsonTrue
	return model.ParseOK
}

func parseFalse(ac *model.AsjsonContext, av *model.AsjsonValue) error {
	if len(ac.JSON) < 5 || ac.JSON[0:5] != "false" {
		av.Typ = model.AsjsonNAT
		return model.ParseInvalidValue
	}
	ac.JSON = ac.JSON[5:]
	av.Typ = model.AsjsonFalse
	return model.ParseOK
}

func parseNumber(ac *model.AsjsonContext, av *model.AsjsonValue) error {
	if len(ac.JSON) == 0 || ac.JSON[0] == '+' || ac.JSON[0] == '.' || ac.JSON[len(ac.JSON)-1] == '.' {
		av.Typ = model.AsjsonNAT
		return model.ParseInvalidValue
	}
	if len(ac.JSON) == 3 {
		invalidNum := strings.ToLower(ac.JSON)
		if invalidNum == "inf" || invalidNum == "nan" {
			av.Typ = model.AsjsonNAT
			return model.ParseInvalidValue
		}
	}
	n, err := strconv.ParseFloat(ac.JSON, 64)
	if err != nil {
		av.Typ = model.AsjsonNAT
		return model.ParseInvalidValue
	}
	av.Typ = model.AsjsonNumber
	av.N = n
	return model.ParseOK
}

func parseValue(ac *model.AsjsonContext, av *model.AsjsonValue) error {
	if len(ac.JSON) == 0 {
		av.Typ = model.AsjsonNAT
		return model.ParseExpectValue
	}
	switch ac.JSON[0] {
	case 'n':
		return parseNull(ac, av)
	case 't':
		return parseTrue(ac, av)
	case 'f':
		return parseFalse(ac, av)
	default:
		return parseNumber(ac, av)
	}
}
