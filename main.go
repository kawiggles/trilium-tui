package main

import(
)

func main() {
	session := StartSession()
	RenderNote(session.GetNote("root"))
}
