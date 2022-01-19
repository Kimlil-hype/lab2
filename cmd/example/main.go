package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	Lab2 "github.com/Kimlil-hype/lab2"
)

func getFlgValue() (inputExp, fileIn, fileOut *string) {
	defer flag.Parse()

	inputExp = flag.String("e", "", "Expression to compute")
	fileIn = flag.String("i", "", "input file")
	fileOut = flag.String("o", "", "output file")
	return
}

func main() {
	var in io.Reader
	var out io.Writer
	var inputExp, fileIn, fileout *string = getFlgValue()
	
	if *inputExp != "" {
		in = strings.NewReader(*inputExp)
	} else if *fileIn != "" {
		in, _ = os.Open(*fileIn)
	}
	
	if *fileIn != "" && *inputExp != "" {
		err := fmt.Errorf("more than one expr is not needed")
		panic(err)
	}

	if *fileout != "" {
		out, _ = os.Create(*fileout)
	} else {
		out = os.Stdout
	}
	handler := Lab2.ComputeHandler{Input: in, Output: out}
	err := handler.Compute()
	if err != nil {
		panic(err)
	}
}
