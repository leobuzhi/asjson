package api

import (
	"fmt"
	"github.com/leobuzhi/asjson/model"
	"github.com/leobuzhi/asjson/parser"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_GetBoolean(t *testing.T) {
	tcs := []struct {
		av  model.AsjsonValue
		b   bool
		err error
	}{
		{
			model.AsjsonValue{Typ: model.AsjsonFalse},
			false,
			nil,
		},
		{
			model.AsjsonValue{Typ: model.AsjsonTrue},
			true,
			nil,
		},
		{
			model.AsjsonValue{Typ: model.AsjsonString},
			false,
			fmt.Errorf("type err,got: %v,want: %v and %v", model.AsjsonString, model.AsjsonTrue, model.AsjsonFalse),
		},
	}

	for _, tc := range tcs {
		b, err := GetBoolean(tc.av)
		assert.Equal(t, tc.b, b)
		assert.Equal(t, tc.err, err)
	}
}

func Test_SetBoolean(t *testing.T) {
	tcs := []struct {
		av  model.AsjsonValue
		typ model.AsjsonType
		b   bool
	}{
		{
			model.AsjsonValue{},
			model.AsjsonFalse,
			false,
		},
		{
			model.AsjsonValue{},
			model.AsjsonTrue,
			true,
		},
	}

	for _, tc := range tcs {
		err := SetBoolean(&tc.av, tc.b)
		assert.Equal(t, nil, err)
		assert.Equal(t, tc.typ, tc.av.Typ)
	}
}

func Test_GetNumber(t *testing.T) {
	tcs := []struct {
		av  model.AsjsonValue
		n   float64
		err error
	}{
		{
			model.AsjsonValue{Typ: model.AsjsonNumber, N: 123},
			123,
			nil,
		},
		{
			model.AsjsonValue{Typ: model.AsjsonNumber, N: 1.23},
			1.23,
			nil,
		},
		{
			model.AsjsonValue{Typ: model.AsjsonString, N: 123},
			0,
			fmt.Errorf("type err,got: %v,want: %v", model.AsjsonString, model.AsjsonNumber),
		},
	}

	for _, tc := range tcs {
		n, err := GetNumber(tc.av)
		assert.Equal(t, tc.n, n)
		assert.Equal(t, tc.err, err)
	}
}

func Test_SetNumber(t *testing.T) {
	tcs := []struct {
		av model.AsjsonValue
		n  float64
	}{
		{
			model.AsjsonValue{},
			123,
		},
	}

	for _, tc := range tcs {
		err := SetNumber(&tc.av, tc.n)
		assert.Equal(t, nil, err)
	}
}

func Test_GetString(t *testing.T) {
	tcs := []struct {
		av  model.AsjsonValue
		s   string
		err error
	}{
		{
			model.AsjsonValue{Typ: model.AsjsonString, S: `"123"`},
			`"123"`,
			nil,
		},
		{
			model.AsjsonValue{Typ: model.AsjsonString, S: `123`},
			`123`,
			nil,
		},
		{
			model.AsjsonValue{Typ: model.AsjsonString, S: "\"1.23\""},
			`"1.23"`,
			nil,
		},
		{
			model.AsjsonValue{Typ: model.AsjsonNumber, S: "123"},
			"",
			fmt.Errorf("type err,got: %v,want: %v", model.AsjsonNumber, model.AsjsonString),
		},
	}

	for _, tc := range tcs {
		s, err := GetString(tc.av)
		assert.Equal(t, tc.s, s)
		assert.Equal(t, tc.err, err)
	}
}

func Test_SetString(t *testing.T) {
	tcs := []struct {
		av model.AsjsonValue
		s  string
	}{
		{
			model.AsjsonValue{},
			"123",
		},
	}

	for _, tc := range tcs {
		err := SetString(&tc.av, tc.s)
		assert.Equal(t, nil, err)
	}
}

func Test_Stringify(t *testing.T) {
	tcs := []struct {
		rawText string
	}{
		{
			"null",
		},
		{
			"  null  ",
		},
		{
			"false",
		},
		{
			"  false  ",
		},
		{
			"true",
		},
		{
			"  true  ",
		},
		{
			" -8.11 ",
		},
		{
			"  2018  ",
		},
		{
			" \"abc\" ",
		},
		{
			"  \"golang is awesome\"  ",
		},
		{
			"[]",
		},
		{
			"[null]",
		},
		{
			"[null,false,true,123,1.23]",
		},
		{
			"[    null   ,  false  ,  true ,   123 ,  1.23]",
		},
		{
			"[[], [   ] , [ 0 ] , [ 0 , 1 ] , [ 0 , 1 , 2 ] ]",
		},
		{
			"[[1,2]] ",
		},
		{
			"[1,[2,3,4,5],4] ",
		},
	}

	for _, tc := range tcs {
		ac := model.AsjsonContext{JSON: tc.rawText}
		var av, newAv model.AsjsonValue
		err := parser.Parse(ac.JSON, &av)
		assert.Equal(t, nil, err)
		avp := &av
		ret, err := Stringify(&avp, av.Len)
		assert.Equal(t, nil, err)
		err = parser.Parse(ret, &newAv)
		assert.Equal(t, nil, err)

		if !reflect.DeepEqual(av, newAv) {
			t.Errorf("Stringify failed,got: %v,want: %v", av, newAv)
		}
	}
}
