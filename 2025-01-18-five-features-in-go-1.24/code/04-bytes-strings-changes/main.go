package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	linesDemo()
	splitSeqDemo()
	splitAfterSeqDemo()
	fieldsSeqDemo()
	fieldsFuncSeq()
}

func fieldsFuncSeq() {
	s := "foo bar+baz asd!"

	f := func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}

	fmt.Println(" In Go 1.23")
	for _, field := range strings.FieldsFunc(s, f) {
		fmt.Println(" ", field)
	}

	fmt.Println(" In Go 1.24")
	for field := range strings.FieldsFuncSeq(s, f) {
		fmt.Println(" ", field)
	}
}

func fieldsSeqDemo() {
	s := " line1     line2 line3   "

	fmt.Println("strings.FieldsSeq")
	fmt.Println(" In Go 1.23")
	for _, line := range strings.Fields(s) {
		fmt.Println("   ", line)
	}

	fmt.Println(" In Go 1.24")
	for line := range strings.FieldsSeq(s) {
		fmt.Println("   ", line)
	}
}

func splitAfterSeqDemo() {
	s := "line 1+line 2+line 3"

	fmt.Println("strings.SplitAfterSeq")
	fmt.Println(" In Go 1.23")
	for _, line := range strings.SplitAfter(s, "+") {
		fmt.Println("   ", line)
	}

	fmt.Println(" In Go 1.24")
	for line := range strings.SplitAfterSeq(s, "+") {
		fmt.Println("   ", line)
	}
}

func splitSeqDemo() {
	s := "line 1+line 2+line 3"

	fmt.Println("strings.SplitSeq")
	fmt.Println(" In Go 1.23")
	for _, line := range strings.Split(s, "+") {
		fmt.Println("   ", line)
	}

	fmt.Println(" In Go 1.24")
	for line := range strings.SplitSeq(s, "+") {
		fmt.Println("   ", line)
	}
}

func linesDemo() {
	s := `line 1
line 2
line 3
`

	fmt.Println("strings.Lines")
	fmt.Println(" In Go 1.23")
	for _, line := range strings.Split(s, "\n") {
		fmt.Println("   ", line)
	}

	fmt.Println(" In Go 1.24")
	for line := range strings.Lines(s) {
		fmt.Print("   ", line)
	}
}
