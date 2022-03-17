package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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
	book1, book2 := Book{}, Book{}

	book1 = book1.SetValues(1, "It", 350, 3, 15, "A125-125-CCD", "1235-4645-1243", "Stephan King")
	book2 = book2.SetValues(2, "White Fang", 424, 4, 18, "A125-122-CCE", "1235-4645-1243", "Jack London")

	var books []Book
	books = append(books, book1)
	books = append(books, book2)

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

//searchCommand() function is called after "search" command, checks and prints if the searched book is in the list.
func searchCommand(books []Book) {
	args := ""

	for i := 1; i < len(flag.Args()); i++ {
		args = args + " " + string(flag.Args()[i])
	}
	args = strings.Title(strings.ToLower(args))[1:]

	flag := []bool{true}
	for _, value := range books {

		if value.Name == args {
			fmt.Printf("\nThe book is found: %s", value.Name)
			flag[0] = false
			break
		}

	}
	if flag[0] {
		fmt.Println("The book is not found!")
	}
}

//readFile() function reads json file is named "data.json", checks for errors then returns data is read from json file
func readFile() map[string]interface{} {

	data := map[string]interface{}{}
	contents, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(contents, &data); err != nil {
		panic(err)
	}

	return data["psychological thriller movies"].(map[string]interface{})
}

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
