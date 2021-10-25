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
	responceProceda.OccurrenceFile.ID = fileProceda.OccurrenceFile.ID
	responceProceda.OccurrenceFile.FileName = fileProceda.OccurrenceFile.FileName
	responceProceda.HeadFile = fileProceda.HeadFile
	responceProceda.HeadFileTwo = fileProceda.HeadFileTwo
	responceProceda.Carrier.CarrierRecordIdentifier = fileProceda.Carrier.CarrierRecordIdentifier
	responceProceda.Carrier.RegisteredNumberCarrier = fileProceda.Carrier.RegisteredNumberCarrier
	responceProceda.Carrier.Name = fileProceda.Carrier.Name
	responceProceda.OccurrenceFile.AmountOccurrences = fileProceda.OccurrenceFile.AmountOccurrences
	responceProceda.OccurrenceFile.AmountRedeployment = fileProceda.OccurrenceFile.AmountRedeployment
}

func setOccurrence(responceProceda *ocoren.OccurrenceProceda,
	fileProceda *ocoren.OccurrenceProceda,
	request *http.Request) {
	perPage, err := strconv.Atoi(request.Header.Get("perPage"))
	checkError(err, "Error: read perPage")
	page, err := strconv.Atoi(request.Header.Get("page"))
	checkError(err, "Error: read page")
	if perPage > fileProceda.OccurrenceFile.AmountOccurrences {
		perPage = fileProceda.OccurrenceFile.AmountOccurrences
	}
	occurrencesInit := (perPage * page) - perPage
	occurrencesEnd := occurrencesInit + perPage
	responceProceda.Carrier.Occurrences = make([]ocoren.Occurrence, 0, perPage)
	for indexOccurrences := occurrencesInit; indexOccurrences < occurrencesEnd; indexOccurrences++ {
		if indexOccurrences < fileProceda.OccurrenceFile.AmountOccurrences {
			responceProceda.Carrier.Occurrences = append(responceProceda.Carrier.Occurrences, fileProceda.Carrier.Occurrences[indexOccurrences])
		}
	}
}
