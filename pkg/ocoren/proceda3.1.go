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

	//TODO: REFACTOR RECORD 342 AND 343. LIST 342 CT-E NOT EXIST
	for line := 0; line < amountLine; line++ {
		recordIdentifier := getRecordIdentifier(originalOcorenSplitLine[line])
		switch recordIdentifier {
		case RECORD_HEARD:
			err = proceda.readHead((originalOcorenSplitLine[line]))
			checkError(err, "Error: to read head")
		case RECORD_HEARD_TWO:
			err = proceda.readHeadTwo(originalOcorenSplitLine[line])
			checkError(err, "Error: to read head two")
		case RECORD_CARRIER:
			err = proceda.carrierDatas(originalOcorenSplitLine[line])
			checkError(err, "Error: to read carrier Datas")
		case RECORD_OCOREN:
			err = proceda.readOccurrences(originalOcorenSplitLine[line], ctePosition, ocorenPosition, occurrences)
			checkError(err, "Error: to read Occurrences")
			proceda.AmountOccurrences++
			ocorenPosition++
		case RECORD_CTE:
			err = proceda.dispacherDatas(originalOcorenSplitLine[line], ctePosition)
			checkError(err, "Error: to read dispacher datas")
			ctePosition = (ctePosition + 1)
			ocorenPosition = 0
			proceda.AmountTransportKnowledges++
		}
	}
	return nil
}

const (
	SENDER_NAME_INIT = 3
	SENDER_NAME_END  = 38

	RECIPIENT_NAME_INIT = 38
	RECIPIENT_NAME_END  = 74

	CREATED_AT_INIT = 73
	CREATED_AT_END  = 83
)

//Cabecalho de intercambio ("000")
func (proceda *OccurrenceProceda) readHead(originalOcorenSplitLine string) (err error) {
	proceda.HeadFileRecordIdentifier = getRecordIdentifier(originalOcorenSplitLine)
	proceda.SenderName = getInformation(originalOcorenSplitLine, SENDER_NAME_INIT, SENDER_NAME_END)
	proceda.RecipientName = getInformation(originalOcorenSplitLine, RECIPIENT_NAME_INIT, RECIPIENT_NAME_END)
	proceda.CreatedAt = getInformation(originalOcorenSplitLine, CREATED_AT_INIT, CREATED_AT_END)
	return nil
}

const (
	FILE_IDENTIFIER_INIT = 3
	FILE_IDENTIFIER_END  = 13
)

//Cabecalho do arquivo("340")
func (proceda *OccurrenceProceda) readHeadTwo(originalOcorenSplitLine string) (err error) {
	proceda.HeadFileTwoRecordIdentifier = getRecordIdentifier(originalOcorenSplitLine)
	proceda.FileIdentifier = getInformation(originalOcorenSplitLine, FILE_IDENTIFIER_INIT, FILE_IDENTIFIER_END)
	return nil
}

const (
	REGISTERED_NUMBER_CARRIER_INIT = 3
	REGISTERED_NUMBER_CARRIER_END  = 17

	CARRIER_NAME_INIT = 17
	CARRIER_NAME_END  = 57

	FILLER_CARRIER_INIT = 57
	FILLER_CARRIER_END  = 119
)

//"341"
func (proceda *OccurrenceProceda) carrierDatas(originalOcorenSplitLine string) (err error) {
	proceda.TransportKnowledges = append(proceda.TransportKnowledges, TransportKnowledge{})
	proceda.CarrierRecordIdentifier = getRecordIdentifier(originalOcorenSplitLine)
	proceda.RegisteredNumberCarrier = getInformation(originalOcorenSplitLine, REGISTERED_NUMBER_CARRIER_INIT, REGISTERED_NUMBER_CARRIER_END)
	proceda.Name = getInformation(originalOcorenSplitLine, CARRIER_NAME_INIT, CARRIER_NAME_END)
	proceda.FillerCarrier = getInformation(originalOcorenSplitLine, FILLER_CARRIER_INIT, FILLER_CARRIER_END)
	return nil
}

const (
	REGISTERED_NUMBER_CTE_INIT = 3
	REGISTERED_NUMBER_CTE_END  = 17

	CONTRACTING_CARRIER_INIT = 17
	CONTRACTING_CARRIER_END  = 27

	SERIES_CTE_INIT = 27
	SERIES_CTE_END  = 32

	NUMBER_CTE_INIT = 32
	NUMBER_CTE_END  = 44
)

//"343" "TransportKnowledge"
func (proceda *OccurrenceProceda) dispacherDatas(originalOcorenSplitLine string, ctePosition int) (err error) {
	proceda.TransportKnowledges = append(proceda.TransportKnowledges, TransportKnowledge{})
	proceda.TransportKnowledges[ctePosition].TransportKnowledgeRecordIdentifier = getRecordIdentifier(originalOcorenSplitLine)
	proceda.TransportKnowledges[ctePosition].RegisteredNumberCte = getInformation(originalOcorenSplitLine, REGISTERED_NUMBER_CTE_INIT, REGISTERED_NUMBER_CTE_END)
	proceda.TransportKnowledges[ctePosition].ContractingCarrier = getInformation(originalOcorenSplitLine, CONTRACTING_CARRIER_INIT, CONTRACTING_CARRIER_END)
	proceda.TransportKnowledges[ctePosition].Series, _ = strconv.Atoi(getInformation(originalOcorenSplitLine, SERIES_CTE_INIT, SERIES_CTE_END))
	proceda.TransportKnowledges[ctePosition].Number, _ = strconv.Atoi(getInformation(originalOcorenSplitLine, NUMBER_CTE_INIT, NUMBER_CTE_END))
	return nil
}

const (
	OCCURRENCE_DATE_INIT = 30
	OCCURRENCE_DATE_END  = 42

	TEXT_INIT = 44
	TEXT_END  = 115

	FILLER_OCCURRENCE_INIT = 115
	FILLER_OCCURRENCE_END  = 119
)

//"342"
func (proceda *OccurrenceProceda) readOccurrences(originalOcorenSplitLine string, ctePosition int, ocorenPosition int, occurrences []OccurrenceCode) (err error) {
	proceda.TransportKnowledges[ctePosition].Occurrences = append(proceda.TransportKnowledges[ctePosition].Occurrences, Occurrence{})
	proceda.TransportKnowledges[ctePosition].Occurrences[ocorenPosition].Invoice = getInvoice(originalOcorenSplitLine)
	proceda.TransportKnowledges[ctePosition].Occurrences[ocorenPosition].OccurrenceCode = getOccurrenceCode(originalOcorenSplitLine, occurrences)
	proceda.TransportKnowledges[ctePosition].Occurrences[ocorenPosition].OccurrenceRecordIdentifier = getRecordIdentifier(originalOcorenSplitLine)
	proceda.TransportKnowledges[ctePosition].Occurrences[ocorenPosition].OccurrenceDate = getInformation(originalOcorenSplitLine, OCCURRENCE_DATE_INIT, OCCURRENCE_DATE_END)
	proceda.TransportKnowledges[ctePosition].Occurrences[ocorenPosition].Text = getInformation(originalOcorenSplitLine, TEXT_INIT, TEXT_END)
	proceda.TransportKnowledges[ctePosition].Occurrences[ocorenPosition].FillerOccurrence = getInformation(originalOcorenSplitLine, FILLER_OCCURRENCE_INIT, FILLER_OCCURRENCE_END)
	proceda.TransportKnowledges[ctePosition].AmountOccurrences++
	return nil
}

const (
	REGISTERED_NUMBER_INVOICE_INIT = 3
	REGISTERED_NUMBER_INVOICE_END  = 17

	SERIES_NFE_INIT = 17
	SERIES_NFE_END  = 20

	NUMBER_NFE_INIT = 20
	NUMBER_NFE_END  = 28
)

func getInvoice(originalOcorenSplitLine string) (invoice Invoice) {
	invoice.RegisteredNumberInvoice = getInformation(originalOcorenSplitLine, REGISTERED_NUMBER_INVOICE_INIT, REGISTERED_NUMBER_INVOICE_END)
	invoice.Series, _ = strconv.Atoi(getInformation(originalOcorenSplitLine, SERIES_NFE_INIT, SERIES_NFE_END))
	invoice.Number, _ = strconv.Atoi(getInformation(originalOcorenSplitLine, NUMBER_NFE_INIT, NUMBER_NFE_INIT))
	return invoice
}

const (
	OCCURRENCE_CODE_INIT = 42
	OCCURRENCE_CODE_END  = 44
)

func getOccurrenceCode(originalOcorenSplitLine string, occurrences []OccurrenceCode) (OccurrenceCode OccurrenceCode) {
	OccurrenceCode.Code, _ = strconv.Atoi(getInformation(originalOcorenSplitLine, OCCURRENCE_CODE_INIT, OCCURRENCE_CODE_END))
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

func getRecordIdentifier(originalOcorenSplitLine string) (recordIdentifier int) {
	information := originalOcorenSplitLine[0:3]
	recordIdentifier, err := strconv.Atoi(information)
	checkError(err, "err: getRecordIdentifier ")
	return recordIdentifier
}

func getInformation(originalOcorenSplitLine string, init int, end int) (information string) {
	information = originalOcorenSplitLine[init:end]
	return information
}

/*
func main() {
	type Ocoren struct {
		A string `init:"25" end:"42"`
	}
	ocoren := Ocoren{}
	st := reflect.TypeOf(ocoren)
	field := st.Field(0)
	init := field.Tag.Get("init")
	end := field.Tag.Get("end")
	fmt.Println(init, end)
}
*/
