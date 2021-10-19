package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SavioAraujoPagung/edi-break-file/pkg/ocoren"
)

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

func setTransportKnowledges(responceProceda *ocoren.OccurrenceProceda,
	fileProceda *ocoren.OccurrenceProceda,
	request *http.Request) {
	maxTransportKnowledges, err := strconv.Atoi(request.Header.Get("maxTransportKnowledges"))
	checkError(err, "Erro: maxTransportKnowledges")
	if maxTransportKnowledges > fileProceda.AmountTransportKnowledges {
		maxTransportKnowledges = fileProceda.AmountTransportKnowledges
	}
	pageTransportKnowledges, err := strconv.Atoi(request.Header.Get("pageTransportKnowledges"))
	checkError(err, "Erro: pageTransportKnowledges")
	transportKnowledgesInit := (maxTransportKnowledges * pageTransportKnowledges) - maxTransportKnowledges
	transportKnowledgesEnd := transportKnowledgesInit + maxTransportKnowledges
	maxOccurrences, err := strconv.Atoi(request.Header.Get("maxOccurrences"))
	checkError(err, "Erro: maxOccurrences")
	pageOccurrences, err := strconv.Atoi(request.Header.Get("pageOccurrences"))
	checkError(err, "Erro: pageOccurrences")
	responceProceda.Carrier.TransportKnowledges = make([]ocoren.TransportKnowledge, 0, maxTransportKnowledges)
	for indexTransportKnowledge := transportKnowledgesInit; indexTransportKnowledge < transportKnowledgesEnd; indexTransportKnowledge++ {
		if indexTransportKnowledge < fileProceda.Carrier.AmountTransportKnowledges {
			responceProceda.Carrier.TransportKnowledges = append(responceProceda.Carrier.TransportKnowledges, fileProceda.Carrier.TransportKnowledges[indexTransportKnowledge])
			indexResponceProceda := len(responceProceda.Carrier.TransportKnowledges) - 1
			setOccurrences(&responceProceda.Carrier.TransportKnowledges[indexResponceProceda],
				&fileProceda.Carrier.TransportKnowledges[indexTransportKnowledge],
				maxOccurrences,
				pageOccurrences)
		}
	}
}

func setOccurrences(responceTransportKnowledge *ocoren.TransportKnowledge, procedaTransportKnowledge *ocoren.TransportKnowledge, maxOccurrences int, pageOccurrences int) {
	if maxOccurrences > procedaTransportKnowledge.AmountOccurrences {
		maxOccurrences = procedaTransportKnowledge.AmountOccurrences
	}
	responceTransportKnowledge.Occurrences = make([]ocoren.Occurrence, 0, maxOccurrences)
	for indexOccurrence := 0; indexOccurrence < maxOccurrences; indexOccurrence++ {
		if indexOccurrence < procedaTransportKnowledge.AmountOccurrences {
			responceTransportKnowledge.Occurrences = append(responceTransportKnowledge.Occurrences, procedaTransportKnowledge.Occurrences[indexOccurrence])
		}
	}
}
