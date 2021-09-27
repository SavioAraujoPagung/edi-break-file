package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/SavioAraujoPagung/edi-break-file/domain"
	"github.com/gorilla/mux"
)

const PORT = ":8080"

func main() {
	muxRoute := mux.NewRouter().StrictSlash(true)
	muxRoute.HandleFunc("/", getProceda).Methods("GET")
	muxRoute.HandleFunc("/proceda", create).Methods("POST")
	fmt.Println("api is online ", PORT)
	log.Fatal(http.ListenAndServe(PORT, muxRoute))
}

type File struct {
	Name string `json: "nome"`
}

type Proceda struct {
	Name              string `json: "nome"`
	AmountOccurrences int    `json: "amount_occurrences"`
}

func create(writer http.ResponseWriter, request *http.Request) {
	var fileProceda domain.OccurrenceProceda

	var file File

	body, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Println("Erro ao ler body")
	}

	err = json.Unmarshal(body, &file)
	if err != nil {
		fmt.Println("Erro no unmarshal")
	}
	fileProceda.FileName = file.Name

	//magica acontece
	err = fileProceda.ReadFile(file.Name)
	if err != nil {
		fmt.Println("Erro ao ler arquivo")
	}

	writer.Header().Set("Content=Type", "application/json")
	json.NewEncoder(writer).Encode(file.Name)
}

func getProceda(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode([]Proceda{{
		Name:              "548245.txt",
		AmountOccurrences: 542,
	}})
}
