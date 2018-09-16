/*
 * @Author: Joey.Chen
 * @Date: 2018-09-10 08:25:48
 * @Last Modified by: Joey.Chen
 * @Last Modified time: 2018-09-15 12:16:44
 */
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/leobuzhi/asjson/api"
	"github.com/leobuzhi/asjson/model"
	"github.com/leobuzhi/asjson/parser"
)

var (
	min = flag.Bool("min", false, "minimize json")
	//for benchmark
	memprofile = flag.String("memprofile", "", "write memory profile to `file`")
	benchmark  bool
)

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

	if benchmark && *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}
