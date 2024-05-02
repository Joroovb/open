package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/joroovb/open/internal/config"
)

var (
	ErrNoExtension     = errors.New("this file has no file extension")
	ErrUnsupportedType = errors.New("unsupported file extension")
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		PrintHelp()
		os.Exit(0)
	}

	pr := config.Get()

	extension, err := GetFileExtension(args[0])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	prog, err := GetProgram(pr, extension, args[0])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	if err := prog.Start(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func GetProgram(pr []config.Program, ext, filepath string) (*exec.Cmd, error) {
	for _, prog := range pr {
		if contains(prog.FileTypes, ext) {
			return exec.Command(prog.Path, filepath), nil
		}
	}
	fmt.Println(ext)
	return nil, ErrUnsupportedType
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetFileExtension(filename string) (string, error) {
	extension := strings.LastIndex(filename, ".")
	if extension == -1 {
		return "", ErrNoExtension
	}
	ext := filename[extension:]
	return ext, nil
}

func PrintHelp() {
	msg := `The universal opening tool.

This application is tool to open many different file formats in with a single command.
Opens the file in a specified program. Can be configured by adding a file extension
and the program that opens the file in config.yml.
		
Usage:
	open <argument>    Opens the file specified in argument.
`

	fmt.Fprint(os.Stdout, msg)
}
