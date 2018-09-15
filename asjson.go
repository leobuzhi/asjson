/*
 * @Author: Joey.Chen
 * @Date: 2018-09-10 08:25:48
 * @Last Modified by: Joey.Chen
 * @Last Modified time: 2018-09-15 00:26:53
 */
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/leobuzhi/asjson/api"
	"github.com/leobuzhi/asjson/model"
	"github.com/leobuzhi/asjson/parser"
)

var min = flag.Bool("min", false, "minimize json")

func main() {
	flag.Parse()

	var av model.AsjsonValue
	var err error
	var ret string

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	err = parser.Parse(string(data), &av)
	if err != nil {
		log.Fatalf("parse json err: %v\n", err)
	}
	head := &av
	if *min {
		ret = api.Stringify(&head)
	} else {
		ret = api.StringBeautify(&head)
	}
	fmt.Println(ret)
}
