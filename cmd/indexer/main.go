package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sakesayso/sakescript"
)

const (
	shouldSortIndex   = false
	defaultIndexLimit = 0
	defaultDirectory  = "community"
)

func main() {
	sortIndex := shouldSortIndex
	indexLimit := defaultIndexLimit
	directory := defaultDirectory

	flag.BoolVar(&sortIndex, "sort", shouldSortIndex, "sort the index by date")
	flag.IntVar(&indexLimit, "limit", defaultIndexLimit, "limit the index to the last n entries")
	flag.StringVar(&directory, "dir", defaultDirectory, "directory to index")
	flag.Parse()

	index, err := sakescript.ParseIndex(directory, directory)
	if err != nil {
		log.Fatal(err)
	}

	// sort by date desc
	if sortIndex {
		index.Sort()
	}

	if indexLimit > 0 && len(index) > indexLimit {
		index = index[len(index)-indexLimit:]
	}

	err = index.Write(directory, fmt.Sprintf("./%s/index.json", directory))
	if err != nil {
		log.Fatal(err)
	}
}
