package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

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
	fmt.Println("File Read:", file.Name)
	fileProceda.FileName = file.Name
	err = fileProceda.ReadFile(file.Name)
	if err != nil {
		fmt.Println("Erro: ler arquivo")
	}
	//writer.Header().Set("Content=Type", "application/json")
	//json.NewEncoder(writer).Encode(fileProceda)
	createResponsePaged(writer, request, &fileProceda)
}

func createResponsePaged(writer http.ResponseWriter, request *http.Request, fileProceda *ocoren.OccurrenceProceda) {
	var responceProceda ocoren.OccurrenceProceda
	setHead(&responceProceda, fileProceda)
	setTransportKnowledges(&responceProceda, fileProceda, request)
	writer.Header().Set("Content=Type", "application/json")
	json.NewEncoder(writer).Encode(responceProceda)
}

func setHead(responceProceda *ocoren.OccurrenceProceda, fileProceda *ocoren.OccurrenceProceda) {
	responceProceda.ID = fileProceda.ID
	responceProceda.FileName = fileProceda.FileName
	responceProceda.HeadFile = fileProceda.HeadFile
	responceProceda.HeadFileTwo = fileProceda.HeadFileTwo
	responceProceda.Carrier.CarrierRecordIdentifier = fileProceda.Carrier.CarrierRecordIdentifier
	responceProceda.Carrier.RegisteredNumberCarrier = fileProceda.Carrier.RegisteredNumberCarrier
	responceProceda.Carrier.Name = fileProceda.Carrier.Name
	responceProceda.AmountOccurrences = fileProceda.AmountOccurrences
	responceProceda.AmountTransportKnowledges = fileProceda.AmountTransportKnowledges
}

func setTransportKnowledges(responceProceda *ocoren.OccurrenceProceda, fileProceda *ocoren.OccurrenceProceda, request *http.Request) {
	maxTransportKnowledges, err := strconv.Atoi(request.Header.Get("maxTransportKnowledges"))
	checkError(err, "Erro: maxTransportKnowledges")
	if maxTransportKnowledges > fileProceda.AmountTransportKnowledges {
		maxTransportKnowledges = fileProceda.AmountTransportKnowledges
	}
	pageTransportKnowledges, err := strconv.Atoi(request.Header.Get("pageTransportKnowledges"))
	checkError(err, "Erro: pageTransportKnowledges")
	transportKnowledgesInit := (maxTransportKnowledges * pageTransportKnowledges) - maxTransportKnowledges
	transportKnowledgesEnd := transportKnowledgesInit + maxTransportKnowledges
	// maxOccurrences, err := strconv.Atoi(request.Header.Get("maxOccurrences"))
	// checkError(err, "Erro: maxOccurrences")
	// pageOccurrences, err := strconv.Atoi(request.Header.Get("pageOccurrences"))
	// checkError(err, "Erro: pageOccurrences")
	responceProceda.Carrier.TransportKnowledges = make([]ocoren.TransportKnowledge, 0, maxTransportKnowledges)
	for indexTransportKnowledge := transportKnowledgesInit; indexTransportKnowledge < transportKnowledgesEnd; indexTransportKnowledge++ {
		if indexTransportKnowledge < fileProceda.Carrier.AmountTransportKnowledges {
			responceProceda.Carrier.TransportKnowledges = append(responceProceda.Carrier.TransportKnowledges, fileProceda.Carrier.TransportKnowledges[indexTransportKnowledge])
			//setOccurrences(responceProceda, fileProceda, maxOccurrences, pageOccurrences)
		}
	}

}

func setOccurrences(responceProceda *ocoren.OccurrenceProceda, fileProceda *ocoren.OccurrenceProceda, maxOccurrences int, pageOccurrences int) {
	// for i := 0; i < maxOccurrences; i++ {
	// 	responceProceda.Carrier.TransportKnowledges[i] = fileProceda.Carrier.TransportKnowledges[i]
	// 	setOccurrences(responceProceda, fileProceda, maxOccurrences, pageOccurrences)
	// }
}

func checkError(err error, message string) {
	if err != nil {
		fmt.Println(message)
	}
}
