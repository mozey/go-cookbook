package main

import (
	"io/ioutil"
	"log"
)

func main() {
	d1 := []byte(`
roses are red
violets are blue
`)
	err := ioutil.WriteFile("poem.txt", d1, 0644)
	if err != nil {
		log.Fatal("Could not generate poem.txt")
	}
}