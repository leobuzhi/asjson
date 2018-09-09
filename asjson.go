package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/leobuzhi/asjson/api"
	"github.com/leobuzhi/asjson/model"
	"github.com/leobuzhi/asjson/parser"
)

var min = flag.Bool("min", false, "minimize json")

func main() {
	flag.Parse()

	s := bufio.NewScanner(os.Stdin)
	var av model.AsjsonValue
	var err error
	var buf bytes.Buffer
	var ret string
	for s.Scan() {
		_, err = buf.Write([]byte(s.Text()))
		if err != nil {
			fmt.Printf("write buffer err: %v\n", err)
		}
	}
	err = parser.Parse(buf.String(), &av)
	if err != nil {
		fmt.Printf("parse json err: %v\n", err)
	}
	head := &av
	if *min {
		ret, err = api.Stringify(&head)
	} else {
		ret, err = api.StringBeautify(&head)
	}

	if err != nil {
		fmt.Printf("stringify json err: %v\n", err)
	}
	fmt.Println(ret)
}
