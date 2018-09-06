package main

import (
	"fmt"

	"github.com/leobuzhi/asjson/api"
	"github.com/leobuzhi/asjson/model"
	"github.com/leobuzhi/asjson/parser"
)

func main() {
	var av model.AsjsonValue
	fmt.Println(parser.Parse("{\"key1\":1,\"key2\":{ \"key3\":3},\"key4\":[1,2,3]}", &av))
	fmt.Println(av)
	for curr := &av; curr != nil; curr = curr.Next {
		fmt.Println(curr.N, curr.S, curr.Typ, curr.Len)
	}
	head := &av
	fmt.Println(api.Stringify(&head))
}
