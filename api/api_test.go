package api

import (
	"reflect"
	"testing"

	"github.com/leobuzhi/asjson/model"
	"github.com/leobuzhi/asjson/parser"
	"github.com/stretchr/testify/assert"
)

func Test_stringify(t *testing.T) {
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
		{
			"{}",
		},
		{
			"{\"key1\":1}",
		},
		{
			"{\"key1\":1,\"key2\":false,\"key3\":[false],\"key4\":{},\"key5\":{\"newkey1\":\"ok\"}}",
		},
	}

	for _, tc := range tcs {
		ac := model.AsjsonContext{JSON: tc.rawText}
		var av, newAv model.AsjsonValue
		err := parser.Parse(ac.JSON, &av)
		assert.Equal(t, nil, err)
		avp := &av
		ret, err := stringify(&avp, av.Len)
		assert.Equal(t, nil, err)
		err = parser.Parse(ret, &newAv)
		assert.Equal(t, nil, err)

		if !reflect.DeepEqual(av, newAv) {
			t.Errorf("stringify failed,got: %v,want: %v", av, newAv)
		}
	}
}

func Test_stringBeautify(t *testing.T) {
	tcs := []struct {
		rawText string
		retText string
	}{
		{
			"null",
			"null",
		},
		{
			"false",
			"false",
		},
		{
			"1.23",
			"1.23",
		},
		{
			"      1.23",
			"1.23",
		},
		{
			"[1,    2,   3]",
			`[
  1,
  2,
  3
]`,
		},
		{
			"[1,    2, [3, 4],  5]",
			`[
  1,
  2,
  [
    3,
    4
  ],
  5
]`,
		},
		{
			`{"key1":  1,"key2":  "2"}`,
			`{
  "key1":1,
  "key2":"2"
}`,
		},
		{
			`{"key1":  "value1",  "key2":  [1 ,2 ,3] , "key3": { "key4":4 }}`,
			`{
  "key1":"value1",
  "key2":[
    1,
    2,
    3
  ],
  "key3":{
    "key4":4
  }
}`,
		},
	}

	for _, tc := range tcs {
		ac := model.AsjsonContext{JSON: tc.rawText}
		var av model.AsjsonValue
		err := parser.Parse(ac.JSON, &av)
		assert.Equal(t, nil, err)
		avp := &av
		ret, err := stringBeautify(&avp, av.Len, 1)
		assert.Equal(t, nil, err)
		assert.Equal(t, tc.retText, ret)
	}
}
