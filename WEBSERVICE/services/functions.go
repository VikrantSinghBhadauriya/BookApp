package services

import (
	"encoding/json"
	"example/webserver/books/data"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Showbooks(w http.ResponseWriter, r *http.Request) {
	bookJSON, err := json.Marshal(&data.Booklist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bookJSON)
}

func ReturnBookById(id int) (*data.Book, int) {
	for i, _ := range data.Booklist {
		if data.Booklist[i].BookId == id {
			return &data.Booklist[i], i
		}

	}
	return nil, -1
}

func ReturnBookbyName(name string) (*data.Book, int) {
	for i, _ := range data.Booklist {
		if data.Booklist[i].Title == name {
			return &data.Booklist[i], i
		}

	}
	return nil, -1
}

func highestID() int {
	var nextId int
	for i, _ := range data.Booklist {
		if nextId < data.Booklist[i].BookId {
			nextId = data.Booklist[i].BookId
		}

	}
	return nextId + 1
}

func ReturnSinglebook(w http.ResponseWriter, r *http.Request) {
	urlsegment := strings.Split(r.URL.Path, "findbook/")
	Bookname := (urlsegment[len(urlsegment)-1])

	fmt.Println(Bookname)
	returnedbook, _ := ReturnBookbyName(Bookname)
	fmt.Println(returnedbook)
	if returnedbook == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	returnedbookJSON, err := json.Marshal(returnedbook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(returnedbookJSON)
}

func Updatebook(w http.ResponseWriter, r *http.Request) {
	var updatedbook data.Book
	urlsegment := strings.Split(r.URL.Path, "updatebook/")

	BId, err := strconv.Atoi(urlsegment[len(urlsegment)-1])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Println(BId)
	_, index := ReturnBookById(BId)

	error := json.Unmarshal([]byte(body), &updatedbook)
	if error != nil {
		log.Fatal(error)
	}

	// fmt.Println(updatedbook)
	data.Booklist[index] = updatedbook
	//fmt.Println(body)
	w.Write(body)
	//fmt.Println(body)

}

func AddNewBook(w http.ResponseWriter, r *http.Request) {
	var newbook data.Book
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	error := json.Unmarshal((body), &newbook)
	if error != nil {
		log.Fatal(error)
	}
	nextbookId := highestID()
	newbook.BookId = nextbookId
	data.Booklist = append(data.Booklist, newbook)
	bookJSON, err := json.Marshal(&data.Booklist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bookJSON)

}
