package model

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles = []Article{
	{Id: 1, Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	{Id: 2, Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}
