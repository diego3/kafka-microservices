package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var customerAPI = "http://localhost:9090/customer"
var httpClient = http.Client{Timeout: time.Second * 5}

type customerPost struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewCustomer(i int) (status string) {
	var payload = customerPost{Name: "test" + strconv.Itoa(i), Email: "test" + strconv.Itoa(i) + "@email.com"}
	content, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("[" + strconv.Itoa(i) + "] error json marshal payload " + err.Error())
		return "fail"
	}
	req, err := http.NewRequest(http.MethodPost, customerAPI, strings.NewReader(string(content)))
	if err != nil {
		log.Println("[" + strconv.Itoa(i) + "] error creating a new request: " + err.Error())
		return "fail"
	}
	req.Header.Add("Content-type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println("[" + strconv.Itoa(i) + "] request error: " + err.Error())
		return "fail"
	}
	return resp.Status
}

func main() {

	var i = 0
	for {
		status := NewCustomer(i)
		log.Println("status: " + status)
		time.Sleep(time.Millisecond * 10)
		i++
	}

}
