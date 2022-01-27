package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type CommandLine struct{}

func (cli *CommandLine) validateArgs() {
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
}

func (cli *CommandLine) proceed(input_file_path, output_file_path string) {

	// Open the file and create the writer
	var f *os.File
	var err error
	if output_file_path == "" {
		f, err = os.Create(fmt.Sprintf("%s_converted.js", input_file_path))
	} else {
		f, err = os.Create(output_file_path)
	}
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	// Open file for reading
	fr, err := os.Open(input_file_path)
	check(err)
	defer fr.Close()

	data, err := ioutil.ReadAll(fr)
	check(err)

	_, err = w.WriteString("let data = new Uint8Array([\n")
	check(err)

	for i := 0; i < len(data)-1; i++ {
		_, err = w.WriteString(fmt.Sprintf("0x%x, ", data[i]))
		check(err)
		if (i+1)%12 == 0 {
			w.WriteString("\n")
		}
	}
	_, err = w.WriteString(fmt.Sprintf("0x%x\n]);", data[len(data)-1]))
	check(err)
	w.Flush()
}

func main() {
	outputFilePtr := flag.String("o", "", "the path to the output file (*.js)")
	cli := CommandLine{}
	flag.Parse()
	cli.validateArgs()

	fmt.Println(*outputFilePtr)

	// get file
	inputFile := flag.Args()[0]
	cli.proceed(inputFile, *outputFilePtr)

	os.Exit(0)
}
