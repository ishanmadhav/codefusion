package main

import "github.com/ishanmadhav/codefusion/pkg/engine"

func main() {
	e := engine.NewEngine()
	err := e.StartEngine()
	if err != nil {
		panic(err)
	}
}
