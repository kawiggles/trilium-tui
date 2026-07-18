package main

import(
)

func main() {
	session := StartSession()
	Render(session.GetNote("root"))
}
