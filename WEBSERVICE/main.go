package main

import (
	"example/webserver/books/data"
	"example/webserver/books/services"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(data.Booklist)
	http.HandleFunc("/showbooks", services.Showbooks)

	http.HandleFunc("/findbook/", services.ReturnSinglebook)
	http.HandleFunc("/updatebook/", services.Updatebook)
	http.HandleFunc("/addbook", services.AddNewBook)

	http.ListenAndServe(":8080", nil)
}
