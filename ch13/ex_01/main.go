package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://dog-facts-api.herokuapp.com/api/v1/resources/dogs?number=1", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("X-My-Client", "Learning Go")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println(res.StatusCode)
}
