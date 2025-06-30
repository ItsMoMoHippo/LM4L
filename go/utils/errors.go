package utils

import (
	"errors"
)

var (
	// arg errors
	ErrNoArgs        = errors.New("no arguments were provided")
	ErrOneArg        = errors.New("only 1 argument was provided, please provide 2 more")
	ErrTwoArgs       = errors.New("only 2 arguments were provided, please provide 1 more")
	ErrTooManyArgs   = errors.New("too many arguments were given")
	ErrDuplicateFile = errors.New("duplicate file inputted, please only provide distinct files")
	// optional flags
	ErrHelp    = errors.New("help")
	ErrVersion = errors.New("version")
)
