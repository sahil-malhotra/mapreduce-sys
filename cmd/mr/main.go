package main

import (
	"fmt"
	"log"
	mr "mapreduce/internal"
	"os"
	"plugin"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Invalid arguments\n")
		os.Exit(1)
	}

	iscoord := os.Args[1] == "mrcoordinator"
	isworker := os.Args[1] == "mrworker"

	if iscoord {
		runCoordinator()
	} else if isworker {
		runWorker()
	} else {
		fmt.Fprintf(os.Stderr, "Invalid arguments\n")
		os.Exit(1)
	}

}

func runWorker() {
	mapf, reducef := loadPlugin(os.Args[2])
	mr.Worker(mapf, reducef)
}

func runCoordinator() {
	m := mr.MakeCoordinator(os.Args[2:], 10)
	for m.Done() == false {
		time.Sleep(time.Second)
	}

	time.Sleep(time.Second)
}

func loadPlugin(filename string) (func(string, string) []mr.KeyValue, func(string, []string) string) {
	p, err := plugin.Open(filename)
	if err != nil {
		log.Fatalf("cannot load plugin %v", filename)
	}
	xmapf, err := p.Lookup("Map")
	if err != nil {
		log.Fatalf("cannot find Map in %v", filename)
	}
	mapf := xmapf.(func(string, string) []mr.KeyValue)
	xreducef, err := p.Lookup("Reduce")
	if err != nil {
		log.Fatalf("cannot find Reduce in %v", filename)
	}
	reducef := xreducef.(func(string, []string) string)

	return mapf, reducef
}
