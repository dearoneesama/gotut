package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var secret = []byte("HelloFromGoTerrible")

func encodeText(text string) string {
	bytes := []byte(text)
	for i := 0; i < len(bytes); i++ {
		bytes[i] = bytes[i] ^ secret[i % len(secret)]
	}
	return base64.StdEncoding.EncodeToString(bytes)
}

func decodeText(text string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}
	for i := 0; i < len(bytes); i++ {
		bytes[i] = bytes[i] ^ secret[i % len(secret)]
	}
	return string(bytes), nil
}

var Info = log.New(os.Stdout, "\u001b[34mINFO: \u001B[0m", log.LstdFlags|log.Lshortfile)
var Warning = log.New(os.Stdout, "\u001b[33mWARNING: \u001B[0m", log.LstdFlags|log.Lshortfile)
var Error = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags | log.Lshortfile)
var Debug = log.New(os.Stdout, "\u001b[36mDEBUG: \u001B[0m", log.LstdFlags|log.Lshortfile)

func handleMain(resp http.ResponseWriter, req * http.Request) {
	if req.Method != "GET" {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	xss, err := decodeText(req.URL.Path[1:])
	if err != nil {
		Error.Println(req, err)
		xss = ""
	}
	t := fmt.Sprintf(
`<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>A Go webpage</title>
  </head>
  <body>
    <h1>A go website</h1>
    <h2>APIs</h2>
    <ul>
      <li>
        <h3>POST /encode</h2>
        Body: <pre>{
  content: string  // string to encode
}</pre>
Result (200): <pre>{
  content: string  // encoded string
}</pre>
      </li>
      <li>
        <h3>POST /decode</h2>
        Body: <pre>{
  content: string  // string to decode
}</pre>
Result (200): <pre>{
  content: string  // decoded string
}</pre>
      </li>
    </ul>
    <h2>Pages</h2>
    <ul>
      <li><h3>/<code>[hidden script]</code></h3> This page</li>
    </ul>
  %v
  </body>
</html>`, xss)
	resp.Header().Add("Content-Type", "text/html; charset=utf-8")
	if _, err := resp.Write([]byte(t)); err != nil {
		Error.Println(req, err)
	}
}

type Payload struct {
	Content string `json:"content"`
}

func handleEncode(resp http.ResponseWriter, req * http.Request) {
	if req.Method != "POST" {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		Error.Println(req, err)
		return
	}
	var posted Payload
	if err := json.Unmarshal(body, &posted); err != nil {
		Error.Println(req, err)
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	if posted.Content == "" {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	resJson, err := json.Marshal(Payload{encodeText(posted.Content)})
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		Error.Println(req, err)
		return
	}
	resp.Header().Add("Content-Type", "application/json")
	if _, err := resp.Write(resJson); err != nil {
		Error.Println(req, err)
	}
}

func handleDecode(resp http.ResponseWriter, req * http.Request) {
	if req.Method != "POST" {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		Error.Println(req, err)
		return
	}
	var posted Payload
	if err := json.Unmarshal(body, &posted); err != nil {
		Error.Println(req, err)
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	if posted.Content == "" {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	decoded, err := decodeText(posted.Content)
	if err != nil {
		Error.Println(req, err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	resJson, err := json.Marshal(Payload{decoded})
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		Error.Println(req, err)
		return
	}
	resp.Header().Add("Content-Type", "application/json")
	if _, err := resp.Write(resJson); err != nil {
		Error.Println(req, err)
	}
}

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/encode", handleEncode)
	http.HandleFunc("/decode", handleDecode)
	Info.Println("Server starting")
	Error.Fatal(http.ListenAndServe(":8766", nil))
}
