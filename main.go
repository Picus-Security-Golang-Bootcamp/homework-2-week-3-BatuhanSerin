package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Book struct {
	ID        int
	Name      string
	Page      int
	Stock     int
	Cost      int
	StockCode string
	ISBN      string
	Author    struct {
		ID   int
		Name string
	}
}

//SetValues set book values
func (b Book) SetValues(id int, name string, page, stock, cost int, stockCode, isbn string, authorId int, authorName string) Book {
	b.ID = id
	b.Name = name
	b.Page = page
	b.Stock = stock
	b.Cost = cost
	b.StockCode = stockCode
	b.ISBN = isbn
	b.Author.ID = authorId
	b.Author.Name = authorName
	return b
}

func main() {

	book1, book2, book3 := Book{}, Book{}, Book{}

	book1 = book1.SetValues(1, "It", 350, 3, 15, "A125-125-CCD", "1235-4645-1243", 2, "Stephan King")
	book2 = book2.SetValues(2, "White Fang", 424, 4, 18, "A125-122-CCE", "1235-4645-1243", 5, "Jack London")
	book3 = book3.SetValues(3, "Harry Potter", 654, 5, 25, "AB13-123-DCE", "1235-4623-1223", 1, "J. K. Rowling")

	var books []Book
	books = append(books, book1)
	books = append(books, book2)
	books = append(books, book3)

	mux := &sync.Mutex{}

	flag.String("com", "ss", "Usage")
	flag.Usage = func() {
		fmt.Println("sssss")
	}
	flag.Parse()

	if flag.Args()[0] == "list" {

		listCommand(books)

	} else if flag.Args()[0] == "search" {

		searchCommand(books)

	} else if flag.Args()[0] == "get" {

		getCommand(books)

	} else if flag.Args()[0] == "delete" {

		deleteCommand(books)

	} else if flag.Args()[0] == "buy" {

		buyCommand(books, mux)

	} else {
		usage()
	}

}

//getValues print elements of given struct
func getValues(value Book) {
	fmt.Printf("\nID: %d, \nName: %s, \nPage %d,\nStock %d,\nCost %d,\nStock Code: %s,\nISBN %s,\nAuthor %d %s,\n",
		value.ID, value.Name, value.Page, value.Stock, value.Cost, value.StockCode, value.ISBN, value.Author.ID, value.Author.Name)
}

//listCommand() function is called after "list" command, prints movie list.
func listCommand(books []Book) {
	fmt.Printf("\n**************Book List**************")

	for _, value := range books {

		getValues(value)
	}
}

//searchCommand is called after "search" command, checks and prints if the searched book is in the list.
func searchCommand(books []Book) {
	if len(flag.Args()) > 1 {

		args := ""

		for i := 1; i < len(flag.Args()); i++ {
			args = args + " " + string(flag.Args()[i])
		}
		args = strings.ToLower(args)[1:]

		flag := []bool{true}
		for _, value := range books {

			if strings.ToLower(value.Name) == args {
				fmt.Printf("\nThe book is found: %s", value.Name)
				getValues(value)
				flag[0] = false
			} else if strings.Contains(strings.ToLower(value.Name), args) {
				fmt.Printf("\n%s includes %s\n", value.Name, args)
				getValues(value)
				flag[0] = false
			}

		}
		if flag[0] {

			fmt.Println("The book is not found!")

		}
	} else {
		fmt.Println("'search' command usage: go run main.go search <bookName>")
	}
}

//getCommand finds book by ID and prints
func getCommand(books []Book) {
	if 3 > len(flag.Args()) && len(flag.Args()) > 1 {

		args, err := strconv.Atoi(flag.Args()[1])

		if err != nil {
			fmt.Println(err)
		}

		flag := []bool{true}
		for _, value := range books {

			if value.ID == args {
				fmt.Printf("\nThe book ID is found: \n")
				getValues(value)
				flag[0] = false
				break
			}

		}
		if flag[0] {
			fmt.Println("The book is not found!")
		}
	} else {
		fmt.Println("'get' command usage: go run main.go get <ID>")
	}
}

//deleteCommand deletes book with ID and prints updated list
func deleteCommand(books []Book) {
	if 3 > len(flag.Args()) && len(flag.Args()) > 1 {
		args, err := strconv.Atoi(flag.Args()[1])

		if err != nil {
			fmt.Println(err)

		}
		flag := []bool{true}
		for key, value := range books {

			if value.ID == args {
				fmt.Printf("\nThe book is deleted(ID:%d): \n", value.ID)

				books = append(books[:key], books[key+1:]...)

				fmt.Printf("\nThe new list: \n")
				listCommand(books)

				flag[0] = false
				break
			}

		}
		if flag[0] {
			fmt.Println("The book is not found!")
		}

	}

}

//buyCommand decrease stock of book has given ID, then prints the book information
func buyCommand(books []Book, mux *sync.Mutex) {
	if 4 > len(flag.Args()) && len(flag.Args()) > 2 {
		id, err := strconv.Atoi(flag.Args()[1])
		quantity, err2 := strconv.Atoi(flag.Args()[2])

		if err != nil || err2 != nil {
			fmt.Println(err, err2)
		}
		flag := []bool{true}
		for _, value := range books {

			if value.ID == id {

				if (value.Stock - quantity) >= 0 {
					mux.Lock()
					value.Stock = value.Stock - quantity
					mux.Unlock()
					getValues(value)
				} else {
					fmt.Printf("There is no enough Stock for ID:%d", value.ID)
				}

				flag[0] = false
				break
			}

		}
		if flag[0] {
			fmt.Println("The book is not found!")
		}

	}

}

//usage prints usage of commnds
func usage() {
	fmt.Printf("Usage:\nto show book list -> 'go run main.go list'\nto search book -> 'go run main.go search <bookName>'\nto get book information with ID -> 'go run main.go get <bookID>'\nto delete book with ID -> 'go run main.go delete <bookID>'\nto buy book with ID -> 'go run main.go buy <bookID> <quantity>'")
}
