package main

import "github.com/ishanmadhav/codefusion/pkg/server"

func main() {
	server := server.NewServer()
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
