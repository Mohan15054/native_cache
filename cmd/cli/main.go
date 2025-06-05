package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Define CLI flags
	action := flag.String("action", "", "Action to perform: set, get, delete")
	key := flag.String("key", "", "Key for the cache operation")
	value := flag.String("value", "", "Value for the set operation")
	ttl := flag.Int("ttl", 60, "TTL for the set operation (in seconds)")
	flag.Parse()

	// Validate input
	if *action == "" || *key == "" {
		fmt.Println("Error: action and key are required")
		flag.Usage()
		return
	}

	// Perform the requested action
	switch *action {
	case "set":
		if *value == "" {
			fmt.Println("Error: value is required for set action")
			return
		}
		setKey(*key, *value, *ttl)
	case "get":
		getKey(*key)
	case "delete":
		deleteKey(*key)
	default:
		fmt.Println("Error: unknown action. Use set, get, or delete")
	}
}

func setKey(key, value string, ttl int) {
	url := fmt.Sprintf("http://localhost:8080/set?key=%s&value=%s&ttl=%d", key, value, ttl)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func getKey(key string) {
	url := fmt.Sprintf("http://localhost:8080/get?key=%s", key)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func deleteKey(key string) {
	url := fmt.Sprintf("http://localhost:8080/delete?key=%s", key)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
