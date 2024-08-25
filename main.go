package main

import (
	"fmt"
	"net/http"
	"log"
	"io"
	"encoding/json"
)

type Fact struct{
	Fact string `json:"fact"`
}

func main(){
	res, err := http.Get("https://catfact.ninja/fact")
	defer res.Body.Close()
	
	if err != nil {
		log.Fatal(err)
	}
	
	body, err := io.ReadAll(res.Body)
	var fact Fact
	err = json.Unmarshal(body, &fact)

	if err != nil{
		log.Fatal(err)
	} 

	fmt.Println(fact.Fact)
}
