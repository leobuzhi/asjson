package parser

import (
	"github.com/leobuzhi/asjson/model"
	"github.com/stretchr/testify/assert"
	"testing"
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
	}
	for _, tc := range tcs {
		var av model.AsjsonValue
		err := parseArray(tc.ac, &av)
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
			&model.AsjsonContext{JSON: "+1.2"},
			model.AsjsonValue{N: 0, Typ: model.AsjsonNAT},
			model.ParseInvalidValue,
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
