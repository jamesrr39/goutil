package main

import (
	"github.com/jamesrr39/goutil/profile/profileviz"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	dataFilePath := kingpin.Arg("data file path", "file path to the profile data").Required().String()
	outFilePath := kingpin.Arg("out file path", "file path to the out file").Required().String()
	kingpin.Parse()

	err := profileviz.Generate(*dataFilePath, *outFilePath)
	if err != nil {
		panic(err)
	}
}
