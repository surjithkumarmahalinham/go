package main

import "fmt"

const max = 3

type Books struct {
	title   string
	author  string
	book_id int
}

func main() {
	var book1 Books
	var book2 Books
	book1.title = "Go"
	book1.author = "Msksurjith"
	book1.book_id = 34534534

	book2.title = "React"
	book2.author = "surjith"
	book2.book_id = 45654445

	fmt.Printf("Book 1 title : %s\n", book1.title)
	fmt.Printf("Book 1 author : %s\n", book1.author)
	fmt.Printf("Book 1 book_id : %d\n", book1.book_id)

	// fmt.Printf("Book 2 title : %s\n", book2.title)
	fmt.Printf("Book 2 author : %s\n", book2.author)
	fmt.Printf("Book 2 book_id : %d\n", book2.book_id)

	PrintBook(&book1)

	PrintBook(&book2)
}
func PrintBook(book *Books) {
	fmt.Printf("Book 2 title : %s\n", book.title)
	fmt.Printf("Book 2 author : %s \n", book.author)
	fmt.Printf("Book 2 book_id : %d \n", book.book_id)
}
