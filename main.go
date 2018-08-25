package main

import (
	"fmt"
	"github.com/leobuzhi/asjson/model"
	"github.com/leobuzhi/asjson/parse"
)

func main() {
	var av model.AsjsonValue
	fmt.Println(parse.Parse("true", &av))
	fmt.Println(av)
}
