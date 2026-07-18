package main

import(
)

func main() {
	session := StartSession()
	ctx := StyleContext{
		Bold: false,
		Italic: false,
		LinkTarget: "",
	}
	RenderNote(session.GetNote("root"))
}
