/*
PROJECT: bf2bf
DESCRIPTION: An interpreter that converts Brainfuck code to Brainfuck code
AUTHOR: Vahin Sharma
*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func compileCommand(command rune) string {
	compiledCode := ""
	switch command {
	case '+', '-', '.', ',', '>', '<', '[', ']':
		compiledCode += string(command)
		
	default:
		compiledCode += "ic" // ic == invalid command
	}

	return compiledCode
}

func main() {
	var inputFile = flag.String("f", "", "Input file")
	var outputFile = flag.String("o", "output.bf", "Output file")

	flag.Parse()

	if *inputFile == "" {
		fmt.Printf("FLAGS:\n")
		fmt.Printf("-f=<file> Input file\n")
		fmt.Printf("-o=<file> Output file (default value: output.bf)\n")
		os.Exit(1)
	}

	code, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("ERROR: Could not read file `%s`\nReason: %s\n", *inputFile, err)
		os.Exit(1)
	}

	var compiledCode string = ``

	for _, c := range code {
		switch c {
		case '\n', '\t', ' ':
			continue
		default:
			commandToAdd := compileCommand(rune(c))
			if commandToAdd != "ic" {
				compiledCode += commandToAdd
			} else {
				fmt.Printf("WARNING: Invalid command `%c`, skipping...\n", c)
			}
		}
	}

	err = ioutil.WriteFile(*outputFile, []byte(compiledCode), 0o644)
	if err != nil {
		fmt.Printf("ERROR: Could not write file `%s`\nReason: %s\n", *outputFile, err)
		os.Exit(2)
	}
}
