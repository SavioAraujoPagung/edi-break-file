package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Proceda struct {
	Name              string `json: "nome"`
	AmountOccurrences int    `json: "amount_occurrences"`
}

func main() {
	resp, err := http.Get("http://localhost:8080/proceda")
	if err != nil {
		log.Fatal()
	}

	if resp.StatusCode != 200 {
		fmt.Println("Error", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)

	var responce []Proceda
	err = json.Unmarshal(body, &responce)

	if err != nil {
		log.Fatal()
	}

	fmt.Println(responce)

}
