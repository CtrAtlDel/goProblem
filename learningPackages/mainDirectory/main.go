package main

import (
	"fmt"
	"learningPackages/cleaning"
	"learningPackages/stuffDirectory"
)

func main() {
	var objWorker = stuffDirectory.Worker{
		Age:  21,
		Name: "Bob",
	}

	var objStuff = stuffDirectory.Stuff{
		Age:  18,
		Name: "John",
	}

	var objCleaning = cleaning.Cleaner{}
	fmt.Println(stuffDirectory.Foo())
	fmt.Println(stuffDirectory.Functionality())
	fmt.Printf("Worker age: %T \n", objWorker.Age)
	fmt.Print("Worker age:\n", objWorker.Age)
	fmt.Println("Worker name: " + string(objWorker.Name))
	fmt.Println("Worker age: ", objStuff.Age)
	fmt.Println("Cleaner name: " + objCleaning.Name)
	fmt.Println("Cleaner age: ", objCleaning.Age)
}
