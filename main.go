package main

import (
	"fmt"
	"github.com/leobuzhi/asjson/model"
	"github.com/leobuzhi/asjson/parser"
)

func main() {
	var av model.AsjsonValue
	fmt.Println(parser.Parse("true", &av))
	fmt.Println(av)
}
