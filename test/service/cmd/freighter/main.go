package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func errorResponse(w http.ResponseWriter, req *http.Request) {

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func testQuery(w http.ResponseWriter, req *http.Request) {
	q := ""
	if req.URL.Query() != nil {
		for key, elements := range req.URL.Query() {
			for _, element := range elements {
				if q != "" {
					q = q + "&"
				}
				q = q + key + "=" + element
			}
		}
	}
	fmt.Fprintf(w, q)
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func testBody(w http.ResponseWriter, req *http.Request) {

	bodyBytes, _ := ioutil.ReadAll(req.Body)
	bodyString := string(bodyBytes)
	fmt.Fprintf(w, bodyString)
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	host := os.Getenv("TEST_HOST")
	gatewayHost := os.Getenv("GATEWAY_HOST")
	values := map[string]string{"serviceName": "test", "host": "http://" + host, "port": "8000"}

	jsonValue, _ := json.Marshal(values)
	resp, _ := http.Post("http://"+gatewayHost+":8080/config", "application/json", bytes.NewBuffer(jsonValue))
	resp.Body.Close()

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/testquery", testQuery)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/error", errorResponse)
	http.HandleFunc("/testbody", testBody)

	http.ListenAndServe(":8000", nil)
}
