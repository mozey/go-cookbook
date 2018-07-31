package main

import (
	"log"
	"io/ioutil"
)

//go:generate go run gen/main.go

func main() {
	poem, err := ioutil.ReadFile("poem.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(poem))
}