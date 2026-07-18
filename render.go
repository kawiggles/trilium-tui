package main

import(
	"fmt"
	"strings"
	"log"
	
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func RenderNote(metadata *Metadata, body string) {
	if metadata == nil {
		log.Printf("Error: no metadata found")
	}
	fmt.Println(metadata.Title)

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

	ctx := StyleContext{
		Bold: false,
		Italic: false,
		LinkTarget: "",
	}

	for _, n := range nodes {
		traverseTree(n, 0, ctx)
	}
}

func traverseTree(n *html.Node, depth int, ctx StyleContext) {
	switch n.Type {
	case html.ElementNode:
		switch n.Data {
		case "strong", "b":
			ctx.Bold = true
		case "em", "i":
			ctx.Italic = true
		case "a":
			ctx.LinkTarget = getTargetId(n)
		}
	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		if text == "" {
			emitSpan(text, ctx)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseTree(c, depth + 1, ctx)
	}
}

func getTargetId(n *html.Node) string {
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			return attr.Val
		}
	}
	return ""
}

type StyleContext struct {
	Bold bool
	Italic bool
	LinkTarget string
}

func emitSpan(text string, ctx StyleContext) {
}
