package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CallRestApi() {
	getAllUrl := "http://localhost:8081/all"
	getByIdUrl := "http://localhost:8081/article/2"

	// postUrl := "http://localhost:8081/"
	// model := m.Article{Id: 3, Title: "Hello three", Desc: "Article Description", Content: "Article Content"}

	callGetMethods(getAllUrl)
	callGetMethods(getByIdUrl)
}

func callGetMethods(getByIdUrl string) {
	resp, err := http.Get(getByIdUrl)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}
