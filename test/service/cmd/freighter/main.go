package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func errorResponse(w http.ResponseWriter, req *http.Request) {

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	values := map[string]string{"serviceName": "test", "host": "127.0.0.1", "port": "8000"}

	jsonValue, _ := json.Marshal(values)

	resp, _ := http.Post("http://127.0.0.1:8080/config", "application/json", bytes.NewBuffer(jsonValue))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/error", errorResponse)

	http.ListenAndServe(":8000", nil)
}
