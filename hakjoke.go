package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Joke struct {
	ID     string
	Joke   string
	Status int
}

func main() {
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")
	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		log.Println("ERROR :(", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		var j Joke
		err = json.Unmarshal(bodyBytes, &j)
		if err != nil {
			log.Println("Error decoding body:", err)
		}
		fmt.Println(j.Joke)
	}
}
