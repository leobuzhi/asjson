package main

import (
	"fmt"

	"github.com/leobuzhi/asjson/api"
	"github.com/leobuzhi/asjson/model"
	"github.com/leobuzhi/asjson/parser"
)

func main() {
	var av model.AsjsonValue
	fmt.Println(parser.Parse("[    1   ,[2,[],3], 6   ]", &av))
	fmt.Println(av)
	head := &av
	fmt.Println(api.Stringify(&head, av.Len))
}
