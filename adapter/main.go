package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId" xml:"userId"`
	ID        int    `json:"id" xml:"id"`
	Title     string `json:"title" xml:"title"`
	Completed bool   `json:"completed" xml:"completed"`
}

type DataInterface interface {
	GetData() (*Todo, error)
}

type RemoteService struct {
	Remote DataInterface
}

func (r *RemoteService) CallRemoteService() (*Todo, error) {
	return r.Remote.GetData()
}

type JSONBackend struct{}

func (jb *JSONBackend) GetData() (*Todo, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

type XMLBackend struct{}

func (xb *XMLBackend) GetData() (*Todo, error) {
	xmlFile := `
<?xml version="1.0" encoding="UTF-8" ?>
<root>
	<userId>1</userId>
	<id>1</id>
	<title>delectus aut autem (XML)</title>	
	<completed>false</completed>
</root>
`
	var todo Todo
	err := xml.Unmarshal([]byte(xmlFile), &todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func main() {
	jsonBackend := &JSONBackend{}
	jsonAdapter := &RemoteService{Remote: jsonBackend}
	tdFromJSON, err := jsonAdapter.CallRemoteService()
	fmt.Printf("From JSON Adapter: %v\t%v\n", tdFromJSON, err)

	xmlBackend := &XMLBackend{}
	xmlAdapter := &RemoteService{Remote: xmlBackend}
	tdFromXML, err := xmlAdapter.CallRemoteService()
	fmt.Printf("From XML Adapter: %v\t%v\n", tdFromXML, err)
}
