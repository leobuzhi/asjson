package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/leobuzhi/asjson/api"
	"github.com/leobuzhi/asjson/model"
	"github.com/leobuzhi/asjson/parser"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var av model.AsjsonValue
	var err error
	var buf bytes.Buffer
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
	ret, err := api.Stringify(&head)
	if err != nil {
		fmt.Printf("stringify json err: %v\n", err)
	}
	fmt.Println(ret)
}
