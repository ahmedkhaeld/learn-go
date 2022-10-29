package main

import (
	"recover/rdir"
)

func main() {
	defer rdir.ReportPanic()
	rdir.ScanDir("go")
}
