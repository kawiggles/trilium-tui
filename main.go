package main

import(
)

func main() {
	session := StartSession()
	Render(session.GetRoot())
}
