package main

import "echolearn/routes"

func main() {
	e := routes.Init()

	e.Logger.Fatal(e.Start(":3000"))
}
