// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name

type author struct{ name string }

// define a book type with a title and author
type book struct {
	title  string
	author author
}

// define a library type to track a list of books
type library map[string][]book

// define addBook to add a book to the library
func (a library) addBook(newBook book) {
	a[newBook.author.name] = append(a[newBook.author.name], newBook)
}

// define a lookupByAuthor function to find books by an author's name
//
func (a library) findAuthor(author_name string) []book {
	return a[author_name]
}
func main() {
	// create a new library
	mainLibrary := library{}

	// add 2 books to the library by the same author
	mainLibrary.addBook(book{
		title:  "Harry Potter",
		author: author{name: "JK Rowling"}})

	mainLibrary.addBook(book{
		title:  "Harry Potter 2",
		author: author{name: "JK Rowling"}})
	// add 1 book to the library by a different author
	mainLibrary.addBook(book{
		title:  "LOTR",
		author: author{name: "JRR Tolkin"}})

	// dump the library with spew
	spew.Dump(mainLibrary)

	// lookup books by known author in the library
	books := mainLibrary.findAuthor("JRR Tolkin")
	// print out the first book's title and its author's name
	if len(books) != 0 {
		b := books[0]
		fmt.Println(b.title, "by", b.author.name)
		spew.Dump(b)
	}

}
