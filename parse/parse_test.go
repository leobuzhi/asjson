package parse

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
	}
	for _, tc := range tcs {
		typ, err := Parse(tc.json)
		assert.Equal(t, tc.typ, typ)
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
		typ, err := parseNull(tc.ac)
		assert.Equal(t, tc.typ, typ)
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
		typ, err := parseTrue(tc.ac)
		assert.Equal(t, tc.typ, typ)
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
		typ, err := parseFalse(tc.ac)
		assert.Equal(t, tc.typ, typ)
		assert.Equal(t, tc.err, err)
	}
}

func Test_parseValue(t *testing.T) {
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
			"n",
			model.AsjsonNAT,
			model.ParseInvalidValue,
		},
		{
			"null",
			model.AsjsonNULL,
			model.ParseOK,
		},
		{
			"abc",
			model.AsjsonNAT,
			model.ParseInvalidValue,
		},
		{
			"true",
			model.AsjsonTrue,
			model.ParseOK,
		},
		{
			"false",
			model.AsjsonFalse,
			model.ParseOK,
		},
		{
			"fa",
			model.AsjsonNAT,
			model.ParseInvalidValue,
		},
	}

	for _, tc := range tcs {
		typ, err := parseValue(&model.AsjsonContext{JSON: tc.json})
		assert.Equal(t, tc.typ, typ)
		assert.Equal(t, tc.err, err)
	}
}
