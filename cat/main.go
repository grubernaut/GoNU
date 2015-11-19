package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"go/scanner"
	"io"
	"os"
	"strconv"
)

var (
	nonBlank    = flag.Bool("b", false, "Number the non-blank output lines, starting at 1")
	nonPrintEnd = flag.Bool("e", false, "Display non-printing characters, and a '$' at the end of each line")
	numLines    = flag.Bool("n", false, "Number the output lines, starting at 1")
	nonPrintTab = flag.Bool("t", false, "Display non-printing characters, and a '^I' as a tab character")
)

var (
	tab   = []byte{' ', ' ', ' ', ' ', ' '}
	space = []byte{' ', ' '}
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: cat [flags] [file ...]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

// Report Errors
func report(err error) int {
	scanner.PrintError(os.Stderr, err)
	return 2
}

func main() {
	os.Exit(realMain())
}

func realMain() int {
	flag.Usage = usage
	flag.Parse()

	exitCode := 0

	for i := 0; i < flag.NArg(); i++ {
		path := flag.Arg(i)
		if err := processFile(path, os.Stdout); err != nil {
			exitCode = report(err)
		}
	}
	return exitCode
}

func processFile(filename string, out io.Writer) error {
	src, err := os.Open(filename)
	defer src.Close()
	if err != nil {
		return err
	}

	res, err := formatFile(src)

	_, err = out.Write(res)

	return err
}

func formatFile(src *os.File) ([]byte, error) {
	scanner := bufio.NewScanner(src)
	var b bytes.Buffer

	lineNum := 0

	for scanner.Scan() {
		line := []byte(scanner.Text())

		if *nonPrintEnd {
			line = appendEnd(line)
		}

		if *nonPrintTab {
			line = appendTab(line)
		}

		if *numLines || *nonBlank {
			lineNum++
			line = numberLine(line, lineNum)
		}

		line = append(line, '\n')
		b.Write(line)
	}
	return b.Bytes(), nil
}

func numberLine(line []byte, lineNum int) []byte {
	if *nonBlank {
		if len(line) == 0 {
			// Skip blank lines
			out := append(tab, ' ')
			out = append(out, space...)
			out = append(out, []byte(line)...)
			return out
		}
	}

	lineByte := []byte(strconv.Itoa(lineNum))
	out := append(tab, lineByte...)
	out = append(out, space...)
	out = append(out, []byte(line)...)
	return out
}

func appendEnd(line []byte) []byte {
	out := append(line, '$')
	return out
}

func appendTab(line []byte) []byte {
	out := []byte{}
	for _, b := range line {
		if b == '\t' {
			out = append(out, []byte("^I")...)
		} else {
			out = append(out, b)
		}
	}
	return out
}
