package parse

import (
	"github.com/leobuzhi/asjson/model"
)

func Parse(json string) (model.AsjsonType, error) {
	var ac model.AsjsonContext
	ac.JSON = json
	parseWhitespace(&ac)
	typ, err := parseValue(&ac)
	parseWhitespace(&ac)
	if len(ac.JSON) > 0 {
		return model.AsjsonNAT, model.ParseRootNotSingular
	}
	return typ, err
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

func parseNull(ac *model.AsjsonContext) (model.AsjsonType, error) {
	if len(ac.JSON) < 4 || ac.JSON[0:4] != "null" {
		return model.AsjsonNAT, model.ParseInvalidValue
	}
	ac.JSON = ac.JSON[4:]
	return model.AsjsonNULL, model.ParseOK
}

func parseTrue(ac *model.AsjsonContext) (model.AsjsonType, error) {
	if len(ac.JSON) < 4 || ac.JSON[0:4] != "true" {
		return model.AsjsonNAT, model.ParseInvalidValue
	}
	ac.JSON = ac.JSON[4:]
	return model.AsjsonTrue, model.ParseOK
}

func parseFalse(ac *model.AsjsonContext) (model.AsjsonType, error) {
	if len(ac.JSON) < 5 || ac.JSON[0:5] != "false" {
		return model.AsjsonNAT, model.ParseInvalidValue
	}
	ac.JSON = ac.JSON[5:]
	return model.AsjsonFalse, model.ParseOK
}

func parseValue(ac *model.AsjsonContext) (model.AsjsonType, error) {
	if len(ac.JSON) == 0 {
		return model.AsjsonNAT, model.ParseExpectValue
	}
	switch ac.JSON[0] {
	case 'n':
		return parseNull(ac)
	case 't':
		return parseTrue(ac)
	case 'f':
		return parseFalse(ac)
	default:
		return model.AsjsonNAT, model.ParseInvalidValue
	}
}
