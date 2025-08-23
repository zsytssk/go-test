package main

import (
	"go-test/plugins"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	println("hello world")
	plugins.Start()
}
