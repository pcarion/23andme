package parser

import (
	"archive/zip"
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Snip description of a snip from 23 and me file
type Snip struct {
	rsid       string
	chromosome string
	position   int64
	genotype   string
}

// Parser parser description
type Parser struct {
	FileName string
	Snips    []Snip
}

func isZipFileName(fileName string) bool {
	return filepath.Ext(fileName) == ".zip"
}

// NewParser ctor
func NewParser(fileName string) *Parser {
	return &Parser{
		FileName: fileName,
	}
}

// based on the extension of the file, read the first entry of a zip archive
// or the content of the file itself (considered to be text file)
func getFileReader(fileName string) (io.Reader, error) {
	if isZipFileName(fileName) {
		zipReader, err := zip.OpenReader(fileName)

		entryFile := zipReader.File[0]
		reader, err := entryFile.Open()
		return reader, err
	}
	reader, err := os.Open(fileName)
	return reader, err
}

// Parse parse the 23 and me file
func (parser *Parser) Parse() error {
	file, err := getFileReader(parser.FileName)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(file)
	parser.Snips = make([]Snip, 100)
	for scanner.Scan() {
		line := scanner.Text()
		// we skip the comment lines
		if line[0] == '#' {
			continue
		}
		// data are seperated by \t
		// columns are:
		// rsid  chromosome      position        genotype
		s := strings.Split(line, "\t")
		// read fields
		rsid := s[0]
		chromosome := s[1]
		position, err := strconv.ParseInt(s[2], 10, 32)
		if err != nil {
			return err
		}
		genotype := s[3]
		parser.Snips = append(parser.Snips, Snip{
			chromosome: chromosome,
			genotype:   genotype,
			position:   position,
			rsid:       rsid,
		})
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
