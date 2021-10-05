package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const url = "http://localhost:7610/greeter/sayhitoserver"

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Please provide your name as argument")
	}
	name := os.Args[1]

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(
		`{"name":"` + name + `"}`,
	)))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	for k, v := range resp.Header {
		fmt.Printf("%v: %v\n", k, v)
	}
	fmt.Println()
	fmt.Println(string(body))
}
