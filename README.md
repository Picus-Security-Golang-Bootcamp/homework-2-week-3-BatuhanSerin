## Book Library

This project made with Golang and contains a list of books.

### Usage

##### 1 - To show the list of books
```go
go run main.go list
```
##### 2 - To search the book in a list
```go
go run main.go search <bookName>
```
The search command is not case sensitive.
###### Example
```go
go run main.go search Requiem For A Dream 
go run main.go search requiem for a dream
```
##### 3 - To get book information with ID
```go
go run main.go get <bookID>
```
##### 4 - To delete book with ID
```go
go run main.go delete <bookID>
```
##### 5 - To buy book with ID
```go
go run main.go buy <bookID> <quantity>
```
