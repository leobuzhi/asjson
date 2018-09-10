/*
 * @Author: Joey.Chen
 * @Date: 2018-09-10 08:25:36
 * @Last Modified by: Joey.Chen
 * @Last Modified time: 2018-09-10 08:26:26
 */
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
	ParseMissCloseBracket        = fmt.Errorf("miss close bracket")
	ParseMissOpenBrace           = fmt.Errorf("miss open brace")
	ParseMissCloseBrace          = fmt.Errorf("miss close brace")
	ParseMissCommaOrCloseBracket = fmt.Errorf("miss comma or close bracket")
	ParseMissCommaOrCloseBrace   = fmt.Errorf("miss comma or close brace")
	ParseMissColon               = fmt.Errorf("miss colon")
)
