package api

import (
	"fmt"
	"github.com/leobuzhi/asjson/model"
)

func GetBoolean(av model.AsjsonValue) (bool, error) {
	switch av.Typ {
	case model.AsjsonTrue:
		return true, nil
	case model.AsjsonFalse:
		return false, nil
	}
	return false, fmt.Errorf("type err,got: %v,want: %v and %v", av.Typ, model.AsjsonTrue, model.AsjsonFalse)
}

func SetBoolean(av *model.AsjsonValue, b bool) error {
	switch b {
	case true:
		av.Typ = model.AsjsonTrue
	case false:
		av.Typ = model.AsjsonFalse
	}
	return nil
}

func GetNumber(av model.AsjsonValue) (float64, error) {
	if av.Typ == model.AsjsonNumber {
		return av.N, nil
	}
	return 0, fmt.Errorf("type err,got: %v,want: %v", av.Typ, model.AsjsonNumber)
}

func SetNumber(av *model.AsjsonValue, n float64) error {
	av.Typ = model.AsjsonNumber
	av.N = n
	return nil
}

func GetString(av model.AsjsonValue) (string, error) {
	if av.Typ == model.AsjsonString {
		return av.S, nil
	}
	return "", fmt.Errorf("type err,got: %v,want: %v", av.Typ, model.AsjsonString)
}

func SetString(av *model.AsjsonValue, s string) error {
	av.Typ = model.AsjsonString
	av.S = s
	return nil
}
