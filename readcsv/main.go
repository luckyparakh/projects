package main

import (
	"flag"
	"readcsv/src"
)

func main() {
	path := flag.String("path", "./quiz.csv", "Path of csv file")
	timer := flag.Int("timer", 10, "Path of csv file")
	flag.Parse()
	c := src.NewCore(*path, *timer)
	c.ReadCSV()
	c.ReadUI()
}
