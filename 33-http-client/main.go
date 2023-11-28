package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Book struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func main() {
	getAllUrl := "http://localhost:8081/book"
	// getByIdUrl := "http://localhost:8081/book/2"

	postUrl := "http://localhost:8081/book"
	model := Book{Name: "Hello eight", Author: "Article 8", Publication: "Article Content"}

	modeln := Book{Name: "Hello eight", Author: "Article 8", Publication: "Article new content"}

	//putUrl := "http://localhost:8081/book/3"

	bookId := "8"

	callPostMethod(postUrl, model)
	callGetMethods(getAllUrl, "")
	callGetMethods(getAllUrl, bookId)

	callPutMethod(getAllUrl, bookId, modeln)
	callGetMethods(getAllUrl, "")
	callGetMethods(getAllUrl, bookId)
}

func callGetMethods(getByIdUrl string, bookId string) {
	if bookId != "" {
		getByIdUrl = getByIdUrl + "/" + bookId
	}
	resp, err := http.Get(getByIdUrl)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer resp.Body.Close()

	printResponse(resp)
}

func printResponse(resp *http.Response) {
	if data, err := ioutil.ReadAll(resp.Body); err == nil {
		fmt.Println(string(data))
		fmt.Println(resp.Status)
	} else {
		log.Println(err.Error())
	}
}

func callPostMethod(url string, model Book) {
	js, _ := json.Marshal(model)
	res, err := http.Post(url, "application/json", bytes.NewReader(js))
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	log.Println("Post Method")
	printResponse(res)
}

func callPutMethod(putUrl string, bookId string, model Book) {
	putUrl = strings.Join([]string{putUrl, bookId}, "/")
	js, _ := json.Marshal(model)
	req, err := http.NewRequest(http.MethodPut, putUrl, bytes.NewReader(js))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")

	if res, err := http.DefaultClient.Do(req); err == nil {
		defer res.Body.Close()

		log.Println("Put Method")
		printResponse(res)
	}
}
