package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	for {
		requestWithClose(client)
		time.Sleep(time.Second * 1)
	}
}

func requestWithClose(client *http.Client) {

	resp, err := client.Get("http://127.0.0.1:8081")

	if err != nil {
		fmt.Printf("error occurred while fetching page, error: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	//c, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatalf("Couldn't parse response body. %+v", err)
	//}

	fmt.Println("ok")
}
