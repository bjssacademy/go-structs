package main

import "fmt"

type Book struct {
	Title     string
	Author    string
	Pages     int
	Published int
}

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
