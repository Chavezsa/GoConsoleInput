package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	EXIT rune = '0' + iota
	SCAN
	STRING
	BYTE
	RUNE
)

const terminator = '#'

const hintContent = "Which type do you want to run?\n" +
	"0). exit\n" +
	"1). fmt.Scan(...)\n" +
	"2). bufio.Reader.ReadString(...)\n" +
	"3). bufio.Reader.ReadByte(...)\n" +
	"4). bufio.Reader.ReadRune()\n" +
	"Please enter 0-4 and press ENTER:"

func main() {
	fmt.Print(hintContent)
	reader := bufio.NewReader(os.Stdin)

	choice, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
		return
	}
	switch choice {
	case STRING:
		ReadString()
	case SCAN:
		ScanOperation()
	case BYTE:
		ReadByte()
	case RUNE:
		ReadRune()
	case EXIT:
		return
	}
}

func ScanOperation() {
	fmt.Print("\nPlease enter some text and press enter: ")
	var content string
	count, err := fmt.Scan(&content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(count)
	fmt.Println(content)
}

func ReadString() {
	fmt.Printf("\n Please enter some text, ends with %c:\n", terminator)
	content, err := bufio.NewReader(os.Stdin).ReadString(terminator)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strings.NewReplacer(string(terminator), "").Replace(content))
}

//alias for uint8
func ReadByte() {
	fmt.Print("\n Please enter a byte:")
	content, err := bufio.NewReader(os.Stdin).ReadByte()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("You entered %q", content)
}

//rune alias for int32
func ReadRune() {
	fmt.Print("\n Please enter a rune:")
	content, _, err := bufio.NewReader(os.Stdin).ReadRune()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("You entered %c", content)
}
