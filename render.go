package main

import(
	"fmt"
	"strings"
	"log"
	
	"golang.org/x/net/html"
)

func Render(s string) {
	fmt.Println(s)

	reader := strings.NewReader(s)

	root, err := html.Parse(reader)
	if err != nil {
		log.Panicf("Error parsing html: %v\n", err)
	}

	traverseTree(root, 0)
}

func traverseTree(n *html.Node, depth int) {
	indent := strings.Repeat("	", depth)

	switch n.Type {
	case html.ElementNode:
		fmt.Printf("%s <%s>\n", indent, n.Data)
	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		if text != "" {
			fmt.Printf("%s %s\n", indent, text)
		}
	}

	for c := n.FirstChild; c != nil; c = n.NextSibling {
		traverseTree(c, depth + 1)
	}
}
