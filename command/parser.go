package command

import (
	"strings"

	"golang.org/x/net/html"
)

type parser struct {
	Source *html.Node
	Data   []string
}

func newParser(source string) (parser, error) {
	reader := strings.NewReader(source)
	doc, err := html.Parse(reader)
	if err != nil {
		return parser{}, err
	}

	p := parser{
		Source: doc,
		Data:   []string{},
	}

	return p, nil
}

func (p *parser) parse() string {
	p.traverse(p.Source)
	return strings.Join(p.Data, "")
}

func (p *parser) traverse(n *html.Node) {
	if n.Type == html.TextNode {
		p.Data = append(p.Data, n.Data)
	} else {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			p.traverse(c)
		}
		switch n.Data {
		case "p":
			p.Data = append(p.Data, "\n")
		}
	}
}
