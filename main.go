package main

import (
	"fmt"
)

func main() {
	session := StartSession()
	fmt.Println(session.GetRoot())
}
