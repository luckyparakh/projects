package linkparser

import (
	"bytes"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type lp struct {
	data []byte
}

type link struct {
	href string
	text string
}

func NewLp(hf []byte) Ilp {
	return &lp{
		data: hf,
	}
}

func (l *lp) LinkParse() (*[]link, error) {
	doc, err := html.Parse(bytes.NewReader(l.data))
	if err != nil {
		log.Printf("error while parsing data %v\n", err)
		return nil, err
	}
	var f func(n *html.Node)
	links := []link{}
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, at := range n.Attr {
				if at.Key == "href" {
					text := findText(n)
					links = append(links, link{
						href: at.Val,
						text: text,
					})
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return &links, nil
}

func findText(n *html.Node) string {
	if n.Type == html.TextNode && n.Data != "" {
		return n.Data
	}
	// Always better to use string builder
	var sb strings.Builder
	// op := ""
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// op += findText(c)
		sb.WriteString(findText(c))
	}
	// return op
	ret := sb.String()
	return strings.Join(strings.Fields(ret), " ")
}
