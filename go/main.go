package main

import (
	"bufio"
	"fmt"
	"lmgo/utils"
	"log"
	"os"
)

const (
	version = "1.0.0"
	appName = "lmgo"
)

func main() {
	// grab args
	args := os.Args[1:]
	filePaths, err := utils.ParseArgs(args)
	if err != nil {
		switch err {
		case utils.ErrHelp:
			fmt.Printf(`%s - List Merging Tool

USAGE:

	%s <input1> <input2> <output>


ARGUMENTS:

	<input1>    First input text file to merge
	<input2>    Second input text file to merge
	<output>	  Output file path


OPTIONS:

	help, --help, --h        Show this help message
	version, --version, --v  Show version information	


DESCRIPTIONS:
      
	Merges two text files by combining all unique files from the files.
	Duplicates are automatically removed
	Empty lines removed
	(future may toggle duplicate removed)

	- Input files must exist and have .txt extension
	- Output file must not exist already
	- All files must be unique
	

EXAMPLES:
			
	%s file1.txt file2.txt merged.txt
	%s --help 
	%s --version
`, appName, appName, appName, appName, appName)
			os.Exit(0)
		case utils.ErrVersion:
			fmt.Printf("%s version %s\n", appName, version)
			os.Exit(0)
		default:
			log.Fatalf("error: %s", err)
		}
	}

	// open files
	in1, err := os.Open(filePaths.Input1)
	if err != nil {
		log.Fatalf("failed to open first file: %v", err)
	}
	defer in1.Close()
	in2, err := os.Open(filePaths.Input2)
	if err != nil {
		log.Fatalf("failed to open second file: %v", err)
	}
	defer in2.Close()

	// make lines
	var inline1 utils.LineArr
	inline1.Seen = make(map[string]bool)
	var out utils.LineArr
	out.Seen = make(map[string]bool)
	scanner := bufio.NewScanner(in1)
	for scanner.Scan() {
		line := scanner.Text()
		inline1.AddLine(line)
	}
	scanner = bufio.NewScanner(in2)
	for scanner.Scan() {
		line := scanner.Text()
		out.AddLine(line)
	}
	for _, line := range inline1.Lines {
		out.AddLine(line)
	}

	// print to new file
	output, err := os.Create(filePaths.Output)
	if err != nil {
		log.Fatalf("failed to make new file: %v", err)
	}
	defer output.Close()
	for _, line := range out.Lines {
		_, err := output.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("Failed during writing of lines: %v", err)
		}
	}
}
