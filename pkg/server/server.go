package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/SavioAraujoPagung/edi-break-file/pkg/ocoren"
	"github.com/SavioAraujoPagung/edi-break-file/pkg/repositories"
)

type File struct {
	Name string `json:"nome"`
}

type Proceda struct {
	Name              string `json:"nome"`
	AmountOccurrences int    `json:"amount_occurrences"`
}

func Teste(writer http.ResponseWriter, request *http.Request) {
	Teste := ocoren.Test{I: 1, Nm: "teste"}
	repositories.Test(&Teste)
	fmt.Println("ola mundo teste")
}

func TestQuery(writer http.ResponseWriter, request *http.Request) {
	var ocorenCode ocoren.OccurrenceCode
	repositories.FindOccurrenceCode(&ocorenCode, 10)
	fmt.Println(ocorenCode)
}

func Create(writer http.ResponseWriter, request *http.Request) {
	var fileProceda ocoren.OccurrenceProceda
	var file File
	body, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Println("Erro: ler body")
	}
	err = json.Unmarshal(body, &file)
	if err != nil {
		fmt.Println("Erro: unmarshal")
	}
	fileProceda.FileName = file.Name
	err = fileProceda.ReadFile(file.Name)
	if err != nil {
		fmt.Println("Erro: ler arquivo")
	}
	createResponsePaged(writer, request, &fileProceda)
}

func checkError(err error, message string) {
	if err != nil {
		fmt.Println(message)
	}
}
