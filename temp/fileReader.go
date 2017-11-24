package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "os"
)

type Categories struct {
    category string
}

type Conversations struct {
    conversation string
}

type Ai struct {
    categories []Categories `json:"categories"`
    conversations []Conversations `json:"conversations"`
}

func (p Ai) toString() string {
    return toJson(p)
}

func toJson(p interface{}) string {
    bytes, err := json.Marshal(p)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    return string(bytes)
}

func main() {
    jsonFile, err := os.Open("./source/ai.json")
    // if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()
    
    raw, err := ioutil.ReadAll(jsonFile)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    var ai Ai
    json.Unmarshal(raw, &ai)
    fmt.Println(ai)
    for i:=0; i< len(ai.conversations); i++ {
        fmt.Println(ai.conversations[i])
    }
}