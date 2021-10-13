package main

import (
	"fmt"
	"log"
	"net/http"

	server "github.com/SavioAraujoPagung/edi-break-file/pkg/server"
	"github.com/gorilla/mux"
)

const PORT = ":1405"

func main() {
	muxRoute := mux.NewRouter().StrictSlash(true)
	muxRoute.HandleFunc("/proceda", server.Create).Methods("POST")
	muxRoute.HandleFunc("/teste", server.Teste).Methods("POST")
	fmt.Println("api is online ", PORT)
	log.Fatal(http.ListenAndServe(PORT, muxRoute))
}
