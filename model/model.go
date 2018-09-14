/*
 * @Author: Joey.Chen
 * @Date: 2018-09-10 08:25:36
 * @Last Modified by: Joey.Chen
 * @Last Modified time: 2018-09-14 22:31:25
 */
package model

import (
	"fmt"
)

//AsjsonType is json type
type AsjsonType int

const (
	//AsjsonNAT is not a type
	AsjsonNAT AsjsonType = iota
	//AsjsonNULL is type null
	AsjsonNULL
	//AsjsonFalse is type false
	AsjsonFalse
	//AsjsonTrue is type true
	AsjsonTrue
	//AsjsonNumber is type number
	AsjsonNumber
	//AsjsonString is type string
	AsjsonString
	//AsjsonArray is type array
	AsjsonArray
	//AsjsonObject is type object
	AsjsonObject
)

//AsjsonValue is json base data structure
type AsjsonValue struct {
	N    float64
	S    string
	Next *AsjsonValue
	Len  int
	Typ  AsjsonType
}

//AsjsonContext is json content
type AsjsonContext struct {
	JSON string
}

var (
	//ParseOK is parse ok
	ParseOK error

	//ErrParseExpectValue is parse error,invalid value
	ErrParseExpectValue = fmt.Errorf("expect value")

	//ErrParseInvalidValue is parse error,invalid value
	ErrParseInvalidValue = fmt.Errorf("invalid value")

	//ErrParseRootNotSingular is parse error,root not singular
	ErrParseRootNotSingular = fmt.Errorf("root not singular")

	//ErrParseMissQuotationMark is parse error,miss quotation mark
	ErrParseMissQuotationMark = fmt.Errorf("miss quotation mark")

	//ErrParseMissOpenBracket is parse error,miss open bracket
	ErrParseMissOpenBracket = fmt.Errorf("miss open bracket")

	//ErrParseMissCloseBracket is parse error,miss close bracket
	ErrParseMissCloseBracket = fmt.Errorf("miss close bracket")

	//ErrParseMissOpenBrace is parse error,miss open brace
	ErrParseMissOpenBrace = fmt.Errorf("miss open brace")

	//ErrParseMissCloseBrace is parse error,miss close brace
	ErrParseMissCloseBrace = fmt.Errorf("miss close brace")

	//ErrParseMissCommaOrCloseBracket is parse error,miss comma or close bracket
	ErrParseMissCommaOrCloseBracket = fmt.Errorf("miss comma or close bracket")

	//ErrParseMissCommaOrCloseBrace is parse error,miss comma or close brace
	ErrParseMissCommaOrCloseBrace = fmt.Errorf("miss comma or close brace")

	//ErrParseMissColon is parse error,miss colon
	ErrParseMissColon = fmt.Errorf("miss colon")
)
