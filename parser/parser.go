package parser

import (
	"bufio"
	"fmt"
	"os"
)

// Parser parser description
type Parser struct {
	fileName string
}

// NewParser ctor
func NewParser(fileName string) *Parser {
	return &Parser{
		fileName: fileName,
	}
}

// Parse parse the 23 and me file
func (parser *Parser) Parse() error {
	file, err := os.Open(parser.fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
