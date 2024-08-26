package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Fact struct {
	Fact string `json:"fact"`
}

func main() {
	res, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var fact Fact
	err = json.Unmarshal(body, &fact)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fact.Fact)
}
