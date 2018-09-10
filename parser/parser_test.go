/*
 * @Author: Joey.Chen
 * @Date: 2018-09-10 08:25:38
 * @Last Modified by: Joey.Chen
 * @Last Modified time: 2018-09-10 08:26:09
 */
package parser

import (
	"testing"

	"github.com/leobuzhi/asjson/model"
	"github.com/stretchr/testify/assert"
)

func Test_Parse(t *testing.T) {
	tcs := []struct {
		json string
		typ  model.AsjsonType
		err  error
	}{
		{
			"",
			model.AsjsonNAT,
			model.ParseExpectValue,
		},
		{
			"null",
			model.AsjsonNULL,
			model.ParseOK,
		},
		{
			"  null",
			model.AsjsonNULL,
			model.ParseOK,
		},
		{
			"\t  null",
			model.AsjsonNULL,
			model.ParseOK,
		},
		{
			"\n  null",
			model.AsjsonNULL,
			model.ParseOK,
		},
		{
			"\r  null",
			model.AsjsonNULL,
			model.ParseOK,
		},
		{
			"\r\n\t  null",
			model.AsjsonNULL,
			model.ParseOK,
		},
		{
			"\r\n\t  null  ",
			model.AsjsonNULL,
			model.ParseOK,
		},
		{
			"\r\n\t  null  abc",
			model.AsjsonNAT,
			model.ParseRootNotSingular,
		},
		{
			"\"",
			model.AsjsonNAT,
			model.ParseMissQuotationMark,
		},
	}
	for _, tc := range tcs {
		var av model.AsjsonValue
		err := Parse(tc.json, &av)
		assert.Equal(t, tc.typ, av.Typ)
		assert.Equal(t, tc.err, err)
	}

}

func Test_parseNull(t *testing.T) {
	tcs := []struct {
		ac  *model.AsjsonContext
		typ model.AsjsonType
		err error
	}{
		{
			&model.AsjsonContext{JSON: "null"},
			model.AsjsonNULL,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{},
			model.AsjsonNAT,
			model.ParseInvalidValue,
		},
		{
			&model.AsjsonContext{JSON: "nul"},
			model.AsjsonNAT,
			model.ParseInvalidValue,
		},
	}

	for _, tc := range tcs {
		var av model.AsjsonValue
		err := parseNull(tc.ac, &av)
		assert.Equal(t, tc.typ, av.Typ)
		assert.Equal(t, tc.err, err)
	}
}

func Test_parseTrue(t *testing.T) {
	tcs := []struct {
		ac  *model.AsjsonContext
		typ model.AsjsonType
		err error
	}{
		{
			&model.AsjsonContext{JSON: "true"},
			model.AsjsonTrue,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{},
			model.AsjsonNAT,
			model.ParseInvalidValue,
		},
		{
			&model.AsjsonContext{JSON: "tr"},
			model.AsjsonNAT,
			model.ParseInvalidValue,
		},
		{
			&model.AsjsonContext{JSON: "True"},
			model.AsjsonNAT,
			model.ParseInvalidValue,
		},
	}

	for _, tc := range tcs {
		var av model.AsjsonValue
		err := parseTrue(tc.ac, &av)
		assert.Equal(t, tc.typ, av.Typ)
		assert.Equal(t, tc.err, err)
	}
}

func Test_parseFalse(t *testing.T) {
	tcs := []struct {
		ac  *model.AsjsonContext
		typ model.AsjsonType
		err error
	}{
		{
			&model.AsjsonContext{JSON: "false"},
			model.AsjsonFalse,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{},
			model.AsjsonNAT,
			model.ParseInvalidValue,
		},
		{
			&model.AsjsonContext{JSON: "tr"},
			model.AsjsonNAT,
			model.ParseInvalidValue,
		},
		{
			&model.AsjsonContext{JSON: "True"},
			model.AsjsonNAT,
			model.ParseInvalidValue,
		},
		{
			&model.AsjsonContext{JSON: "False"},
			model.AsjsonNAT,
			model.ParseInvalidValue,
		},
	}

	for _, tc := range tcs {
		var av model.AsjsonValue
		err := parseFalse(tc.ac, &av)
		assert.Equal(t, tc.typ, av.Typ)
		assert.Equal(t, tc.err, err)
	}
}

func Test_parseNumber(t *testing.T) {
	tcs := []struct {
		ac  *model.AsjsonContext
		typ model.AsjsonType
		n   float64
		err error
	}{
		{
			&model.AsjsonContext{JSON: "0.0"},
			model.AsjsonNumber,
			0,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "-0"},
			model.AsjsonNumber,
			0,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "-0.0"},
			model.AsjsonNumber,
			0,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "1"},
			model.AsjsonNumber,
			1.0,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "-1"},
			model.AsjsonNumber,
			-1.0,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "3.1416"},
			model.AsjsonNumber,
			3.1416,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "1E10"},
			model.AsjsonNumber,
			1E10,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "1e10"},
			model.AsjsonNumber,
			1e10,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "1E+10"},
			model.AsjsonNumber,
			1E+10,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "-1E10"},
			model.AsjsonNumber,
			-1E10,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "-1E10"},
			model.AsjsonNumber,
			-1E10,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "-1e10"},
			model.AsjsonNumber,
			-1e10,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "-1E+10"},
			model.AsjsonNumber,
			-1E+10,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "-1E-10"},
			model.AsjsonNumber,
			-1E-10,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "1.234E+10"},
			model.AsjsonNumber,
			1.234E+10,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "1.234E-10"},
			model.AsjsonNumber,
			1.234E-10,
			model.ParseOK,
		},
		{
			//note(joey.chen): underflow
			&model.AsjsonContext{JSON: "1e-10000"},
			model.AsjsonNumber,
			0.0,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "+0"},
			model.AsjsonNAT,
			0.0,
			model.ParseInvalidValue,
		},
		{
			&model.AsjsonContext{JSON: "+1"},
			model.AsjsonNAT,
			0.0,
			model.ParseInvalidValue,
		},
		{
			&model.AsjsonContext{JSON: ".123"},
			model.AsjsonNAT,
			0.0,
			model.ParseInvalidValue,
		},
		{
			&model.AsjsonContext{JSON: "1."},
			model.AsjsonNAT,
			0.0,
			model.ParseInvalidValue,
		},
		{
			&model.AsjsonContext{JSON: "INF"},
			model.AsjsonNAT,
			0.0,
			model.ParseInvalidValue,
		},
		{
			&model.AsjsonContext{JSON: "inf"},
			model.AsjsonNAT,
			0.0,
			model.ParseInvalidValue,
		},
		{
			&model.AsjsonContext{JSON: "NAN"},
			model.AsjsonNAT,
			0.0,
			model.ParseInvalidValue,
		},
		{
			&model.AsjsonContext{JSON: "nan"},
			model.AsjsonNAT,
			0.0,
			model.ParseInvalidValue,
		},
		{
			&model.AsjsonContext{JSON: "1E012"},
			model.AsjsonNumber,
			1E012,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "1.0000000000000002"},
			model.AsjsonNumber,
			1.0000000000000002,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "4.9406564584124654e-324"},
			model.AsjsonNumber,
			4.9406564584124654e-324,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "-4.9406564584124654e-324"},
			model.AsjsonNumber,
			-4.9406564584124654e-324,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "2.2250738585072009e-308"},
			model.AsjsonNumber,
			2.2250738585072009e-308,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "-2.2250738585072009e-308"},
			model.AsjsonNumber,
			-2.2250738585072009e-308,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "2.2250738585072014e-308"},
			model.AsjsonNumber,
			2.2250738585072014e-308,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "-2.2250738585072014e-308"},
			model.AsjsonNumber,
			-2.2250738585072014e-308,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "1.7976931348623157e+308"},
			model.AsjsonNumber,
			1.7976931348623157e+308,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "-1.7976931348623157e+308"},
			model.AsjsonNumber,
			-1.7976931348623157e+308,
			model.ParseOK,
		},
	}
	for _, tc := range tcs {
		var av model.AsjsonValue
		err := parseNumber(tc.ac, &av)
		assert.Equal(t, tc.typ, av.Typ)
		assert.Equal(t, tc.n, av.N)
		assert.Equal(t, tc.err, err)
	}
}

func Test_parseString(t *testing.T) {
	tcs := []struct {
		ac  *model.AsjsonContext
		typ model.AsjsonType
		s   string
		err error
	}{
		{
			&model.AsjsonContext{JSON: "\"abc\""},
			model.AsjsonString,
			"abc",
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: `""`},
			model.AsjsonString,
			``,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: `"123"`},
			model.AsjsonString,
			`123`,
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: `"`},
			model.AsjsonNAT,
			``,
			model.ParseMissQuotationMark,
		},
		{
			&model.AsjsonContext{JSON: "\"\""},
			model.AsjsonString,
			"",
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "\"Hello\""},
			model.AsjsonString,
			"Hello",
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "\"Hello\nWorld\""},
			model.AsjsonString,
			"Hello\nWorld",
			model.ParseOK,
		},
	}
	for _, tc := range tcs {
		var av model.AsjsonValue
		err := parseString(tc.ac, &av)
		assert.Equal(t, tc.typ, av.Typ)
		assert.Equal(t, tc.s, av.S)
		assert.Equal(t, tc.err, err)
	}
}

func Test_parseArray(t *testing.T) {
	tcs := []struct {
		ac  *model.AsjsonContext
		typ model.AsjsonType
		len int
		err error
	}{
		{
			&model.AsjsonContext{JSON: "[]"},
			model.AsjsonArray,
			0,
			nil,
		},
		{
			&model.AsjsonContext{JSON: "[null]"},
			model.AsjsonArray,
			1,
			nil,
		},
		{
			&model.AsjsonContext{JSON: "[null"},
			model.AsjsonNAT,
			0,
			model.ParseMissCloseBracket,
		},
		{
			&model.AsjsonContext{JSON: "[null,false,true,123,1.23]"},
			model.AsjsonArray,
			5,
			nil,
		},
		{
			&model.AsjsonContext{JSON: "[    null   ,  false  ,  true ,   123 ,  1.23]"},
			model.AsjsonArray,
			5,
			nil,
		},
		{
			&model.AsjsonContext{JSON: "[1.23,]"},
			model.AsjsonNAT,
			0,
			model.ParseInvalidValue,
		},
		{
			&model.AsjsonContext{JSON: "[    null   ,  false  ,  true ,   123 ,  1.23,  \"123\",\"abc\"]"},
			model.AsjsonArray,
			7,
			nil,
		},
		{
			&model.AsjsonContext{JSON: "[    null   , [1,2], false  ,  true ,   123 ,  1.23,  \"123\",\"abc\"]"},
			model.AsjsonArray,
			8,
			nil,
		},
		{
			&model.AsjsonContext{JSON: "[[], [   ] , [ 0 ] , [ 0 , 1 ] , [ 0 , 1 , 2 ] ]"},
			model.AsjsonArray,
			5,
			nil,
		},
	}
	for _, tc := range tcs {
		var av model.AsjsonValue
		err := parseArray(tc.ac, &av)
		assert.Equal(t, tc.typ, av.Typ)
		assert.Equal(t, tc.len, av.Len)
		assert.Equal(t, tc.err, err)
	}
}

func Test_parseObject(t *testing.T) {
	tcs := []struct {
		ac  *model.AsjsonContext
		typ model.AsjsonType
		len int
		err error
	}{
		{
			&model.AsjsonContext{JSON: "{}"},
			model.AsjsonObject,
			0,
			nil,
		},
		{
			&model.AsjsonContext{JSON: "{ \"key1\": 1    }"},
			model.AsjsonObject,
			2,
			nil,
		},
		{
			&model.AsjsonContext{JSON: "{ \"key2\":[ 1 ] }"},
			model.AsjsonObject,
			2,
			nil,
		},
		{
			&model.AsjsonContext{JSON: "{ \"key2\":[ 1 ] "},
			model.AsjsonNAT,
			0,
			model.ParseMissCloseBrace,
		},
		{
			&model.AsjsonContext{JSON: "{ \"key1\": 1 ,\"key2\":[ 1 ] , \"key3\":\"3\" ,\"key4\": false,\"key5\" : true,\"key6\":null }"},
			model.AsjsonObject,
			12,
			nil,
		},
		{
			&model.AsjsonContext{JSON: "{   :[ 1 ] "},
			model.AsjsonNAT,
			0,
			model.ParseMissQuotationMark,
		},
		{
			&model.AsjsonContext{JSON: "{  1:1, "},
			model.AsjsonNAT,
			0,
			model.ParseMissQuotationMark,
		},
		{
			&model.AsjsonContext{JSON: "{true:1, "},
			model.AsjsonNAT,
			0,
			model.ParseMissQuotationMark,
		},
		{
			&model.AsjsonContext{JSON: "{false:1, "},
			model.AsjsonNAT,
			0,
			model.ParseMissQuotationMark,
		},
		{
			&model.AsjsonContext{JSON: "{null:1, "},
			model.AsjsonNAT,
			0,
			model.ParseMissQuotationMark,
		},
		{
			&model.AsjsonContext{JSON: "{[]:1, "},
			model.AsjsonNAT,
			0,
			model.ParseMissQuotationMark,
		},
		{
			&model.AsjsonContext{JSON: "{{}:1, "},
			model.AsjsonNAT,
			0,
			model.ParseMissQuotationMark,
		},
		{
			&model.AsjsonContext{JSON: "{\"a\":1, "},
			model.AsjsonNAT,
			0,
			model.ParseMissQuotationMark,
		},
		{
			&model.AsjsonContext{JSON: "{\"a\" "},
			model.AsjsonNAT,
			0,
			model.ParseMissColon,
		},
		{
			&model.AsjsonContext{JSON: "{\"a\" ,\"hello \""},
			model.AsjsonNAT,
			0,
			model.ParseMissColon,
		},
		{
			&model.AsjsonContext{JSON: "{\"a\":1"},
			model.AsjsonNAT,
			0,
			model.ParseMissCloseBrace,
		},
		{
			&model.AsjsonContext{JSON: "{\"a\":1]"},
			model.AsjsonNAT,
			0,
			model.ParseMissCommaOrCloseBrace,
		},
		{
			&model.AsjsonContext{JSON: "{\"a\":1 \"b\""},
			model.AsjsonNAT,
			0,
			model.ParseMissCommaOrCloseBrace,
		},
		{
			&model.AsjsonContext{JSON: "{\"a\":{}"},
			model.AsjsonNAT,
			0,
			model.ParseMissCloseBrace,
		},
		{
			&model.AsjsonContext{JSON: "{ \"key1\": 1 ,\"key2\":2 ,\"key3\":{\"key3\": 3 ,\"key4\":4}  }"},
			model.AsjsonObject,
			6,
			nil,
		},
	}
	for _, tc := range tcs {
		var av model.AsjsonValue
		err := parseObject(tc.ac, &av)
		assert.Equal(t, tc.typ, av.Typ)
		assert.Equal(t, tc.len, av.Len)
		assert.Equal(t, tc.err, err)
	}
}

func Test_parseValue(t *testing.T) {
	tcs := []struct {
		ac  *model.AsjsonContext
		av  model.AsjsonValue
		err error
	}{
		{
			&model.AsjsonContext{JSON: ""},
			model.AsjsonValue{N: 0, Typ: model.AsjsonNAT},
			model.ParseExpectValue,
		},
		{
			&model.AsjsonContext{JSON: "n"},
			model.AsjsonValue{N: 0, Typ: model.AsjsonNAT},
			model.ParseInvalidValue,
		},
		{
			&model.AsjsonContext{JSON: "null"},
			model.AsjsonValue{N: 0, Typ: model.AsjsonNULL},
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "abc"},
			model.AsjsonValue{N: 0, Typ: model.AsjsonNAT},
			model.ParseInvalidValue,
		},
		{
			&model.AsjsonContext{JSON: "true"},
			model.AsjsonValue{N: 0, Typ: model.AsjsonTrue},
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "false"},
			model.AsjsonValue{N: 0, Typ: model.AsjsonFalse},
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "fa"},
			model.AsjsonValue{N: 0, Typ: model.AsjsonNAT},
			model.ParseInvalidValue,
		},
		{
			&model.AsjsonContext{JSON: "1.2"},
			model.AsjsonValue{N: 1.2, Typ: model.AsjsonNumber},
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "[1.2]"},
			model.AsjsonValue{N: 0, Typ: model.AsjsonArray},
			model.ParseOK,
		},
		{
			&model.AsjsonContext{JSON: "{\"1\":1.2}"},
			model.AsjsonValue{N: 0, Typ: model.AsjsonObject},
			model.ParseOK,
		},
	}

	for _, tc := range tcs {
		var av model.AsjsonValue
		err := parseValue(tc.ac, &av)
		assert.Equal(t, tc.av.Typ, av.Typ)
		assert.Equal(t, tc.av.N, av.N)
		assert.Equal(t, tc.err, err)
	}
}
