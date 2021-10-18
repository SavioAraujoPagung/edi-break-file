package ocoren

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
)

const RECORD_HEARD = 0
const RECORD_HEARD_TWO = 340
const RECORD_CARRIER = 341
const RECORD_OCOREN = 342
const RECORD_CTE = 343

//Cabecalho de intercambio
const AMOUNT_RECORD_000_BY_FILE = 1

//Cabecalho do arquivo
const AMOUNT_RECORD_340_BY_000 = 1

//Dados da transportadora
const AMOUNT_RECORD_341_BY_340 = 1

//Ocorrencia na entrega
const AMOUNT_RECORD_342_BY_341 = 5000

//Ocorrencia na entrega
const AMOUNT_RECORD_343_BY_342 = 1

//read all content file - OCOREN PROCEDA 3.1
func (proceda *OccurrenceProceda) ReadFile(fileName string, occurrences []OccurrenceCode) (err error) {
	proceda.ID = rand.Intn(100000)
	fileOcoren, err := ioutil.ReadFile(fileName)
	checkError(err, "Error: open file")
	proceda.ContentFile = string(fileOcoren)
	originalOcorenSplitLine := strings.Split(proceda.ContentFile, "\n")
	var ctePosition int = 0
	var ocorenPosition int = 0
	var amountLine int = len(originalOcorenSplitLine)
	proceda.TransportKnowledges = make([]TransportKnowledge, 0, 5000)
	for line := 0; line < amountLine; line++ {
		originalOcorenSplitChar := strings.Split(originalOcorenSplitLine[line], "")
		recordIdentifier := getRecordIdentifier(originalOcorenSplitChar)
		switch recordIdentifier {
		case RECORD_HEARD:
			err = proceda.readHead(originalOcorenSplitChar)
			checkError(err, "Error: to read head")
		case RECORD_HEARD_TWO:
			err = proceda.readHeadTwo(originalOcorenSplitChar)
			checkError(err, "Error: to read head two")
		case RECORD_CARRIER:
			err = proceda.carrierDatas(originalOcorenSplitChar)
			checkError(err, "Error: to read carrier Datas")
		case RECORD_OCOREN:
			err = proceda.readOccurrences(originalOcorenSplitChar, ctePosition, ocorenPosition, occurrences)
			checkError(err, "Error: to read Occurrences")
			proceda.AmountOccurrences++
			ocorenPosition++
		case RECORD_CTE:
			err = proceda.dispacherDatas(originalOcorenSplitChar, ctePosition)
			checkError(err, "Error: to read dispacher datas")
			ctePosition = (ctePosition + 1)
			ocorenPosition = 0
			proceda.AmountTransportKnowledges++
		}
	}

	return nil
}

const SENDER_NAME_INIT = 3
const SENDER_NAME_END = 38

const RECIPIENT_NAME_INIT = 38
const RECIPIENT_NAME_END = 73

const CREATED_AT_INIT = 0
const CREATED_AT_END = 0

//Cabecalho de intercambio ("000")
func (proceda *OccurrenceProceda) readHead(originalOcorenSplitChar []string) (err error) {
	proceda.HeadFileRecordIdentifier = getRecordIdentifier(originalOcorenSplitChar)
	proceda.SenderName = getInformation(originalOcorenSplitChar, SENDER_NAME_INIT, SENDER_NAME_END)
	proceda.RecipientName = getInformation(originalOcorenSplitChar, RECIPIENT_NAME_INIT, RECIPIENT_NAME_END)
	proceda.CreatedAt = getInformation(originalOcorenSplitChar, CREATED_AT_INIT, CREATED_AT_END)
	return nil
}

const FILE_IDENTIFIER_INIT = 3
const FILE_IDENTIFIER_END = 13

//Cabecalho do arquivo("340")
func (proceda *OccurrenceProceda) readHeadTwo(originalOcorenSplitChar []string) (err error) {
	proceda.HeadFileTwoRecordIdentifier = getRecordIdentifier(originalOcorenSplitChar)
	proceda.FileIdentifier = getInformation(originalOcorenSplitChar, FILE_IDENTIFIER_INIT, FILE_IDENTIFIER_END)
	return nil
}

const REGISTERED_NUMBER_CARRIER_INIT = 3
const REGISTERED_NUMBER_CARRIER_END = 17
const CARRIER_NAME_INIT = 17
const CARRIER_NAME_END = 57
const FILLER_CARRIER_INIT = 57
const FILLER_CARRIER_END = 119

//"341"
func (proceda *OccurrenceProceda) carrierDatas(originalOcorenSplitChar []string) (err error) {
	proceda.TransportKnowledges = append(proceda.TransportKnowledges, TransportKnowledge{})
	proceda.CarrierRecordIdentifier = getRecordIdentifier(originalOcorenSplitChar)
	proceda.RegisteredNumberCarrier = getInformation(originalOcorenSplitChar, REGISTERED_NUMBER_CARRIER_INIT, REGISTERED_NUMBER_CARRIER_END)
	proceda.Name = getInformation(originalOcorenSplitChar, CARRIER_NAME_INIT, CARRIER_NAME_END)
	proceda.FillerCarrier = getInformation(originalOcorenSplitChar, FILLER_CARRIER_INIT, FILLER_CARRIER_END)
	return nil
}

const REGISTERED_NUMBER_CTE_INIT = 3
const REGISTERED_NUMBER_CTE_END = 17
const CONTRACTING_CARRIER_INIT = 17
const CONTRACTING_CARRIER_END = 27
const SERIES_CTE_INIT = 27
const SERIES_CTE_END = 32
const NUMBER_CTE_INIT = 32
const NUMBER_CTE_END = 44

//"343" "TransportKnowledge"
func (proceda *OccurrenceProceda) dispacherDatas(originalOcorenSplitChar []string, ctePosition int) (err error) {
	proceda.TransportKnowledges = append(proceda.TransportKnowledges, TransportKnowledge{})
	proceda.TransportKnowledges[ctePosition].TransportKnowledgeRecordIdentifier = getRecordIdentifier(originalOcorenSplitChar)
	proceda.TransportKnowledges[ctePosition].RegisteredNumberCte = getInformation(originalOcorenSplitChar, REGISTERED_NUMBER_CTE_INIT, REGISTERED_NUMBER_CTE_END)
	proceda.TransportKnowledges[ctePosition].ContractingCarrier = getInformation(originalOcorenSplitChar, CONTRACTING_CARRIER_INIT, CONTRACTING_CARRIER_END)
	proceda.TransportKnowledges[ctePosition].Series, _ = strconv.Atoi(getInformation(originalOcorenSplitChar, SERIES_CTE_INIT, SERIES_CTE_END))
	proceda.TransportKnowledges[ctePosition].Number, _ = strconv.Atoi(getInformation(originalOcorenSplitChar, NUMBER_CTE_INIT, NUMBER_CTE_END))
	return nil
}

const OCCURRENCE_DATE_INIT = 30
const OCCURRENCE_DATE_END = 42

const TEXT_INIT = 44
const TEXT_END = 115

const FILLER_OCCURRENCE_INIT = 115
const FILLER_OCCURRENCE_END = 119

//"342"
func (proceda *OccurrenceProceda) readOccurrences(originalOcorenSplitChar []string, ctePosition int, ocorenPosition int, occurrences []OccurrenceCode) (err error) {
	proceda.TransportKnowledges[ctePosition].Occurrences = append(proceda.TransportKnowledges[ctePosition].Occurrences, Occurrence{})
	proceda.TransportKnowledges[ctePosition].Occurrences[ocorenPosition].Invoice = getInvoice(originalOcorenSplitChar)
	proceda.TransportKnowledges[ctePosition].Occurrences[ocorenPosition].OccurrenceCode = getOccurrenceCode(originalOcorenSplitChar, occurrences)
	proceda.TransportKnowledges[ctePosition].Occurrences[ocorenPosition].OccurrenceRecordIdentifier = getRecordIdentifier(originalOcorenSplitChar)
	proceda.TransportKnowledges[ctePosition].Occurrences[ocorenPosition].OccurrenceDate = getInformation(originalOcorenSplitChar, OCCURRENCE_DATE_INIT, OCCURRENCE_DATE_END)
	proceda.TransportKnowledges[ctePosition].Occurrences[ocorenPosition].Text = getInformation(originalOcorenSplitChar, TEXT_INIT, TEXT_END)
	proceda.TransportKnowledges[ctePosition].Occurrences[ocorenPosition].FillerOccurrence = getInformation(originalOcorenSplitChar, FILLER_OCCURRENCE_INIT, FILLER_OCCURRENCE_END)
	proceda.TransportKnowledges[ctePosition].AmountOccurrences++
	return nil
}

const REGISTERED_NUMBER_INVOICE_INIT = 3
const REGISTERED_NUMBER_INVOICE_END = 17

const SERIES_NFE_INIT = 17
const SERIES_NFE_END = 20

const NUMBER_NFE_INIT = 20
const NUMBER_NFE_END = 28

func getInvoice(originalOcorenSplitChar []string) (invoice Invoice) {
	invoice.RegisteredNumberInvoice = getInformation(originalOcorenSplitChar, REGISTERED_NUMBER_INVOICE_INIT, REGISTERED_NUMBER_INVOICE_END)
	invoice.Series, _ = strconv.Atoi(getInformation(originalOcorenSplitChar, SERIES_NFE_INIT, SERIES_NFE_END))
	invoice.Number, _ = strconv.Atoi(getInformation(originalOcorenSplitChar, NUMBER_NFE_INIT, NUMBER_NFE_INIT))
	return invoice
}

const OCCURRENCE_CODE_INIT = 42
const OCCURRENCE_CODE_END = 44

func getOccurrenceCode(originalOcorenSplitChar []string, occurrences []OccurrenceCode) (OccurrenceCode OccurrenceCode) {
	OccurrenceCode.Code, _ = strconv.Atoi(getInformation(originalOcorenSplitChar, OCCURRENCE_CODE_INIT, OCCURRENCE_CODE_END))
	OccurrenceCode = occurrences[OccurrenceCode.Code]
	return OccurrenceCode
}

func checkError(err error, message string) {
	if err != nil {
		fmt.Println(message)
		fmt.Println(err)
		//panic(err)
	}
}

func getRecordIdentifier(originalOcorenSplitChar []string) (recordIdentifier int) {
	record := originalOcorenSplitChar[0] + originalOcorenSplitChar[1] + originalOcorenSplitChar[2]
	recordIdentifier, err := strconv.Atoi(record)
	checkError(err, "err: getRecordIdentifier "+record)
	return recordIdentifier
}

func getInformation(originalOcorenSplitChar []string, init int, end int) (information string) {
	for i := init; i < end; i++ {
		information = information + originalOcorenSplitChar[i]
	}
	return information
}
