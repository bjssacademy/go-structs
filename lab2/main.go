package main

import "fmt"

type Book struct {
	Title     string
	Author    string
	Pages     int
	Published int
}

//create a new struct "Price" which contains 
//Cost as an integer
//Currency as a string 
//Now add the Price struct to the Book struct
//Don't forget to initialise it!
//Print out the Cost and Currency

func main() {
	
	goBook := Book{
		Title:     "The Go Programming Language",
		Author:    "Alan A. A. Donovan & Brian W. Kernighan",
		Pages:     380,
		Published: 2015,
	}

	fmt.Println("Title:", goBook.Title)
	fmt.Println("Author:", goBook.Author)
	fmt.Println("Pages:", goBook.Pages)
	fmt.Println("Published:", goBook.Published)
}
