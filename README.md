# 23andme
Parser for 23 and me export file

This program _only_ parses the file and makes available the SNPs in a data structure, as an array of

```go
type Snip struct {
	rsid       string
	chromosome string
	position   int64
	genotype   string
}
```

This could be the starting point to further process those data.

A starting point could be to look at the projects listed there:

https://github.com/plashchynski/awesome-genetics


# How to use

- download your raw data from 23andme and store the file in the `data` directory.
- This file can either be a acompressed file (`.zip` extension) or a text file.

- then run:

```
go run cmd/23andme/main.go --data <path to your data file>
```

