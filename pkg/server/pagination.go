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
	setOccurrence(&responceProceda, fileProceda, request)
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
	responceProceda.AmountRedeployment = fileProceda.AmountRedeployment
}

func setOccurrence(responceProceda *ocoren.OccurrenceProceda,
	fileProceda *ocoren.OccurrenceProceda,
	request *http.Request) {
	perPage, err := strconv.Atoi(request.Header.Get("perPage"))
	checkError(err, "Error: read perPage")
	page, err := strconv.Atoi(request.Header.Get("page"))
	checkError(err, "Error: read page")
	if perPage > fileProceda.AmountOccurrences {
		perPage = fileProceda.AmountOccurrences
	}
	occurrencesInit := (perPage * page) - perPage
	occurrencesEnd := occurrencesInit + perPage
	responceProceda.Carrier.Occurrences = make([]ocoren.Occurrence, 0, perPage)
	for indexOccurrences := occurrencesInit; indexOccurrences < occurrencesEnd; indexOccurrences++ {
		if indexOccurrences < fileProceda.AmountOccurrences {
			responceProceda.Carrier.Occurrences = append(responceProceda.Carrier.Occurrences, fileProceda.Carrier.Occurrences[indexOccurrences])
		}
	}
}
