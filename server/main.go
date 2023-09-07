package main

import (
	"fmt"

	"github.com/poboisvert/GO_Fiber/server/server"
)

type Todo struct {
	ID    int  `json:"id"`
	Title int  `json:"title"`
	Done  bool `json:"done"`
	Body  bool `json:"body"`
}

func main() {
	fmt.Print("Hello World")
	server.InitializeApp()
}
