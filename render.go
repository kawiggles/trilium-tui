package main

import(
	"fmt"
	"strings"
	"log"
	
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func Render(metadata string, body string) {
	fmt.Println(metadata)
	fmt.Println(body)

	reader := strings.NewReader(body)
	context := &html.Node{
		Type: html.ElementNode,
		Data: "body",
		DataAtom: atom.Body,
	}

	nodes, err := html.ParseFragment(reader, context)
	if err != nil {
		log.Printf("Error parsing html: %v\n", err)
		return
	}

	for _, n := range nodes {
		traverseTree(n, 0)
	}
}

func traverseTree(n *html.Node, depth int) {
	indent := strings.Repeat("\t", depth)

	switch n.Type {
	case html.ElementNode:
		fmt.Printf("%s<%s>\n", indent, n.Data)
	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		if text != "" {
			fmt.Printf("%s%s\n", indent, text)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseTree(c, depth + 1)
	}
}
