package main

import (
	"fmt"
	"log"

	"github.com/pcarion/23andme/parser"

	"github.com/segmentio/cli"
)

func main() {
	type config struct {
		Datafile string `flag:"-d,--data" help:"23 and me export file (txt)"`
	}

	cli.Exec(cli.Command(func(config config) {
		fmt.Printf("data file: %s\n", config.Datafile)
		p := parser.NewParser(config.Datafile)
		err := p.Parse()
		if err != nil {
			log.Fatalf("error parsing file: %s", config.Datafile)
		}
		fmt.Printf("Number of SNPs: %d\n", len(p.Snips))
	}))
}
