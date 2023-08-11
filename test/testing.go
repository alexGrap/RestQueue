package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	models "inter/internal"
	"inter/internal/middleware"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	path, _ := os.Getwd()
	file, _ := os.ReadFile(path + "/test/example.json")
	body := models.Task{}
	err := json.Unmarshal(file, &body)
	if err != nil {
		log.Fatal(err)
	}
	//testingPost(body)
	client := http.Client{}
	testingPost(&client, body)
}

func testingPost(client *http.Client, body models.Task) {
	postCounter := 0
	errCounter := 0
	js, _ := json.Marshal(body)
	bodyReader := bytes.NewReader(js)
	req, _ := http.NewRequest(
		"POST", "http://127.0.0.1:3000/create", bodyReader)
	fmt.Println("\n\nTest 1st. Valid posting/getting")
	resp, err := client.Do(req)

	if resp.Status == "201 Created" {
		postCounter++
		fmt.Printf("Posting: Done: %d, err: %d\n", postCounter, errCounter)
	} else {
		errCounter++
		fmt.Printf("Posting: Done: %d, err: %d\n", postCounter, errCounter)
		fmt.Println(err)
	}
	testingGet(client, 1)

	resp, err = client.Do(req)
	if resp.Status == "201 Created" {
		postCounter++
		fmt.Printf("Posting: Done: %d, err: %d\n", postCounter, errCounter)
	} else {
		errCounter++
		fmt.Printf("Posting: Done: %d, err: %d\n", postCounter, errCounter)
		fmt.Println(err)
	}
	testingGet(client, 2)

	fmt.Println("\n\nTest posting unvalid")
	postBody := models.Task{StartElement: 12, Delta: 15, L: 5, TTL: 15}
	js, _ = json.Marshal(postBody)
	bodyReader = bytes.NewReader(js)
	req, _ = http.NewRequest("POST", "http://127.0.0.1:3000/create", bodyReader)
	resp, err = client.Do(req)
	if resp.Status == "400 Bad Request" {
		postCounter++
	} else {
		errCounter++
		fmt.Println(resp.Status, err)
	}
	fmt.Printf("Posting: Done: %d, err: %d\n", postCounter, errCounter)

	postBody = models.Task{ElementCount: 12, StartElement: 15, L: 5, TTL: 15}
	js, _ = json.Marshal(postBody)
	bodyReader = bytes.NewReader(js)
	req, _ = http.NewRequest("POST", "http://127.0.0.1:3000/create", bodyReader)
	resp, err = client.Do(req)
	if resp.Status == "400 Bad Request" {
		postCounter++
	} else {
		errCounter++
		fmt.Println(resp.Status, err)
	}
	fmt.Printf("Posting: Done: %d, err: %d\n", postCounter, errCounter)

	postBody = models.Task{ElementCount: 12, StartElement: 15, Delta: 5, TTL: 15}
	js, _ = json.Marshal(postBody)
	bodyReader = bytes.NewReader(js)
	req, _ = http.NewRequest("POST", "http://127.0.0.1:3000/create", bodyReader)
	resp, err = client.Do(req)
	if resp.Status == "400 Bad Request" {
		postCounter++
	} else {
		errCounter++
		fmt.Println(resp.Status, err)
	}
	fmt.Printf("Posting: Done: %d, err: %d\n", postCounter, errCounter)

	postBody = models.Task{ElementCount: 12, StartElement: 15, Delta: 5, L: 15}
	js, _ = json.Marshal(postBody)
	bodyReader = bytes.NewReader(js)
	req, _ = http.NewRequest("POST", "http://127.0.0.1:3000/create", bodyReader)
	resp, err = client.Do(req)
	if resp.Status == "400 Bad Request" {
		postCounter++
	} else {
		errCounter++
		fmt.Println(resp.Status, err)
	}
	fmt.Printf("Posting: Done: %d, err: %d\n", postCounter, errCounter)

}

func testingGet(client *http.Client, expected int) {
	var body []models.Task
	getCounter := expected - 1
	errCounter := 0
	req, _ := http.NewRequest(
		"GET", "http://127.0.0.1:3000/get", nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	json.Unmarshal(resBody, &body)
	for i, _ := range body {
		if !middleware.Validation(body[i]) {
			errCounter++
			fmt.Printf("Getting: Done: %d, err: %d\n", getCounter, errCounter)
			fmt.Println("uncorrected body parsing")
			return
		}
	}

	if len(body) == expected && expected == 1 {
		getCounter++
		fmt.Printf("Getting: Done: %d, err: %d\n", getCounter, errCounter)

	} else if len(body) == expected && expected == 2 {
		if body[0].Status == "On going" && body[1].Status == "On going" {
			getCounter++
			fmt.Printf("Getting: Done: %d, err: %d\n", getCounter, errCounter)
			time.Sleep(10 * time.Second)
			resp, _ = client.Do(req)
			resBody, _ = io.ReadAll(resp.Body)
			json.Unmarshal(resBody, &body)
			if body[0].Status == "Done" && body[1].Status == "Done" {
				getCounter++
				fmt.Printf("Getting: Done: %d, err: %d\n", getCounter, errCounter)

			} else {
				errCounter++
				fmt.Printf("Getting: Done: %d, err: %d\n", getCounter, errCounter)
				fmt.Println("uncorrected body parsing")
			}

		} else {
			errCounter++
			fmt.Printf("Getting: Done: %d, err: %d\n", getCounter, errCounter)
			fmt.Println("uncorrected body parsing")
		}

	} else {
		errCounter++
		fmt.Printf("Getting: Done: %d, err: %d\n", getCounter, errCounter)
		fmt.Println("failed get multi")
	}
}
