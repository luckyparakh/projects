package main

import (
	"flag"
	"io/ioutil"
	linkparser "linkparser/src"
	"log"
	"os"
)

func main() {
	hf := flag.String("hf", "./example/ex1.html", "Path to html file to parse")
	flag.Parse()
	if *hf != "" {
		f, e := readFile(*hf)
		if e != nil {
			log.Fatalf("Error while opening file %v\n", e)
		}
		linkP := linkparser.NewLp(f)
		links, err := linkP.LinkParse()
		if err != nil {
			log.Fatalf("Error while parsing links: %v\n", err)
		} else {
			log.Println(links)
		}
	}

}
func readFile(f string) ([]byte, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}
