package data

type Book struct {
	BookId  int    `json:"bookid"`
	Title   string `json:"tile"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Booklist []Book
var BookJSONN []byte

func init() {

	BookJSONN = []byte(`{"bookid":2,"tile":"book1","desc":"dfghjk ","content":"Article"}`)
	Booklist = []Book{
		{BookId: 1, Title: "book2", Desc: "Article Description", Content: "Article Content"},
		{BookId: 2, Title: "NewBook", Desc: "Article Description", Content: "Article Content"},
	}
}
