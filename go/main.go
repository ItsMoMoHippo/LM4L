package main

import (
	"fmt"
	"lmgo/utils"
	"log"
	"os"
)

func main() {
	// grab args
	args := os.Args[1:]
	for _, arg := range args {
		fmt.Println(arg)
	}
	filePaths, err := utils.ParseArgs(args)
	if err != nil {
		switch err {
		case utils.ErrHelp:
			fmt.Println("foo")
			os.Exit(0)
		case utils.ErrVersion:
			fmt.Println("bar")
			os.Exit(0)
		default:
			log.Fatalf("error: %s", err)
		}
	}
	_ = filePaths
}
