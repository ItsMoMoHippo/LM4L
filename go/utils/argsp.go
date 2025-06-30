package utils

import (
	"fmt"
	"os"
	"strings"
)

type ArgFiles struct {
	input1 string
	input2 string
	output string
}

func ParseArgs(args []string) (ArgFiles, error) {
	// help and version flag
	if len(args) == 1 {
		switch args[0] {
		case "version", "--version", "--v":
			return ArgFiles{}, ErrVersion
		case "help", "--help", "--h":
			return ArgFiles{}, ErrHelp
		}
	}

	// bad arg amounts
	switch len(args) {
	case 0:
		return ArgFiles{}, ErrNoArgs
	case 1:
		return ArgFiles{}, ErrOneArg
	case 2:
		return ArgFiles{}, ErrTwoArgs
	case 3:
	default:
		return ArgFiles{}, ErrTooManyArgs
	}

	// check for txt
	for _, file := range args {
		x := checkExt(file)
		if !x {
			return ArgFiles{}, fmt.Errorf("file '%s' is not a text file, please only use text files", file)
		}
	}

	// check for dumplicates
	if args[0] == args[1] || args[0] == args[2] || args[1] == args[2] {
		return ArgFiles{}, ErrDuplicateFile
	}

	fileCheck, err := fileExists(args[0])
	if err != nil {
		return ArgFiles{}, err
	}
	if !fileCheck {
		return ArgFiles{}, fmt.Errorf("file '%s' does not exist, please enter an existing file for input", args[0])
	}
	fileCheck, err = fileExists(args[1])
	if err != nil {
		return ArgFiles{}, err
	}
	if !fileCheck {
		return ArgFiles{}, fmt.Errorf("file '%s' does not exist, please enter an existing file for input", args[1])
	}
	fileCheck, err = fileExists(args[2])
	if err != nil {
		return ArgFiles{}, err
	}
	if fileCheck {
		return ArgFiles{}, fmt.Errorf("file '%s' does exist already, please choose a different name for the output", args[2])
	}

	// return successfully
	return ArgFiles{
		input1: args[0],
		input2: args[1],
		output: args[2],
	}, nil
}

func checkExt(file string) bool {
	if len(file) < 4 {
		return false
	}

	_, ext, found := strings.Cut(file, ".")

	if !found {
		return false
	}

	if ext != "txt" {
		return false
	}

	return true
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
