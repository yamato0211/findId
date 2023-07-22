package main

import (
	"github.com/yamato0211/findId"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(findId.Analyzer) }
