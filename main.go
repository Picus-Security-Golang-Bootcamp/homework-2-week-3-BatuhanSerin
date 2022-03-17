<<<<<<< HEAD
package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Book struct {
	ID        int
	Name      string
	Page      int
	Stock     int
	Cost      int
	StockCode string
	ISBN      string
	Author    string
}

func (b Book) SetValues(id int, name string, page, stock, cost int, stockCode, isbn, author string) Book {
	b.ID = id
	b.Name = name
	b.Page = page
	b.Stock = stock
	b.Cost = cost
	b.StockCode = stockCode
	b.ISBN = isbn
	b.Author = author
	return b
}

func main() {
	book1, book2, book3 := Book{}, Book{}, Book{}

	book1 = book1.SetValues(1, "It", 350, 3, 15, "A125-125-CCD", "1235-4645-1243", "Stephan King")
	book2 = book2.SetValues(2, "White Fang", 424, 4, 18, "A125-122-CCE", "1235-4645-1243", "Jack London")
	book3 = book3.SetValues(3, "Harry Potter", 654, 5, 25, "AB13-123-DCE", "1235-4623-1223", "J. K. Rowling")

	var books []Book
	books = append(books, book1)
	books = append(books, book2)
	books = append(books, book3)

	//names := readFile()

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

	}

}

//listCommand() function is called after "list" command, prints movie list.
func listCommand(books []Book) {
	fmt.Printf("\n**************Book List**************")

	for _, value := range books {

		fmt.Printf("\nID: %d, \nName: %s, \nPage %d,\nStock %d,\nCost %d,\nStock Code: %s,\nISBN %s,\nAuthor %s,\n",
			value.ID, value.Name, value.Page, value.Stock, value.Cost, value.StockCode, value.ISBN, value.Author)

	}
}

//searchCommand is called after "search" command, checks and prints if the searched book is in the list.
func searchCommand(books []Book) {
	if len(flag.Args()) > 1 {

		args := ""

		for i := 1; i < len(flag.Args()); i++ {
			args = args + " " + string(flag.Args()[i])
		}
		args = strings.Title(strings.ToLower(args))[1:]

		flag := []bool{true}
		for _, value := range books {

			if value.Name == args {
				fmt.Printf("\nThe book is found: %s", value.Name)
				fmt.Printf("\nID: %d, \nName: %s, \nPage %d,\nStock %d,\nCost %d,\nStock Code: %s,\nISBN %s,\nAuthor %s,\n",
					value.ID, value.Name, value.Page, value.Stock, value.Cost, value.StockCode, value.ISBN, value.Author)
				flag[0] = false
				break
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
				fmt.Printf("\nID: %d, \nName: %s, \nPage %d,\nStock %d,\nCost %d,\nStock Code: %s,\nISBN %s,\nAuthor %s,\n",
					value.ID, value.Name, value.Page, value.Stock, value.Cost, value.StockCode, value.ISBN, value.Author)
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

//readFile() function reads json file is named "data.json", checks for errors then returns data is read from json file
// func readFile() map[string]interface{} {

// 	data := map[string]interface{}{}
// 	contents, err := ioutil.ReadFile("data.json")
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := json.Unmarshal(contents, &data); err != nil {
// 		panic(err)
// 	}

// 	return data["psychological thriller movies"].(map[string]interface{})
// }

// //listCommand() function is called after "list" command, prints movie list.
// func listCommand(names map[string]interface{}) {
// 	fmt.Println("**************Movie List**************")
// 	for _, movieNames := range names {

// 		movieNames := fmt.Sprintf("%v", movieNames)

// 		fmt.Println(movieNames)
// 	}
// }

//searchCommand() function is called after "search" command, checks and prints if the searched movie is in the list.
// func searchCommand(names map[string]interface{}) {
// 	args := ""

// 	for i := 2; i < len(os.Args); i++ {
// 		args = args + " " + string(os.Args[i])
// 	}
// 	args = strings.Title(strings.ToLower(args))[1:]

// 	flag := []bool{true}
// 	for _, movieNames := range names {

// 		movieNames := fmt.Sprintf("%v", movieNames)

// 		if movieNames == args {
// 			fmt.Printf("The movie is found: %s", movieNames)
// 			flag[0] = false
// 			break
// 		}

// 	}
// 	if flag[0] {
// 		fmt.Println("The movie is not found!")
// 	}
// }
