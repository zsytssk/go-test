package main

import (
	"os"
	"path/filepath"
)

type Candidate struct {
	Text   string
	Weight float64
	Score  int
}

func main() {
	ex, err := os.Executable()
	if err != nil {
		return
	}
	basename := filepath.Base(ex)
	println(basename)
}
