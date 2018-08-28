package model

import (
	"fmt"
)

type AsjsonType int

const (
	AsjsonNAT AsjsonType = iota
	AsjsonNULL
	AsjsonFalse
	AsjsonTrue
	AsjsonNumber
	AsjsonString
	AsjsonArray
	AsjsonObject
)

type AsjsonValue struct {
	N    float64
	S    string
	Next *AsjsonValue
	Len  int
	Typ  AsjsonType
}

type AsjsonContext struct {
	JSON string
}

var (
	ParseOK                      error
	ParseExpectValue             = fmt.Errorf("expect value")
	ParseInvalidValue            = fmt.Errorf("invalid value")
	ParseRootNotSingular         = fmt.Errorf("root not singular")
	ParseMissQuotationMark       = fmt.Errorf("miss quotation mark")
	ParseMissOpenBracket         = fmt.Errorf("miss open bracket")
	ParseMissCommaOrCloseBracket = fmt.Errorf("miss comma or close bracket")
)
