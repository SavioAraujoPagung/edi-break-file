package domain

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
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

//Cabeçalho do arquivo - "000"
type HeadFile struct {
	HeadFileRecordIdentifier int       `json: "identificador"`
	SenderName               string    `json: "remetente"`
	RecipientName            string    `json: "destinatario"`
	CreatedAt                time.Time `json: "data_criacao"`
	Filler                   string    `json: -`
}

//Cabeçalho dois - "340"
type HeadFileTwo struct {
	HeadFileTwoRecordIdentifier int    `json: "identificador"`
	FileIdentifier              string `json: "identificador_arquivo"`
	Filler                      string `json: -`
}

//Informação de transportadora - "341"
type Carrier struct {
	CarrierRecordIdentifier int    `json: "identificador"`
	RegisteredNumber        string `json: "cnpj_transportadora"`
	Name                    string `json: "nome_transportadora"`
	Filler                  string `json: "-"`
	TransportKnowledges     []TransportKnowledge
}

//Conhecimento de transporte CT-e - "343"
type TransportKnowledge struct {
	TransportKnowledgeRecordIdentifier int    `json: "identificador"`
	RegisteredNumber                   string `json: "cgc_contratante"`
	ContractingCarrier                 string `json: "transportadora_contratante"`
	Series                             int    `json: "cte_serie"`
	Number                             int    `json: "cte_numero"`
	Occurrences                        []Occurrence
}

//Nota fiscal - NF-e
type Invoice struct {
	RegisteredNumber string `json: "nfe_cnpj_emitente"`
	Series           int    `json: "nfe_serie"`
	Number           int    `json: "nfe_numero"`
}

//Codigo da ocorrencia - vide tabela de ocorrencias Proceda-3.1
type OccurrenceCode struct {
	Code        int    `json: "codigo_ocorrencia"`
	Description string `json: "nome_ocorrencia"`
}

//Informações sobre uma ocorrencia - "342"
type Occurrence struct {
	Invoice                    []Invoice
	OccurrenceCode             []OccurrenceCode
	OccurrenceRecordIdentifier int       `json: "identificador"`
	OccurrenceDate             time.Time `json: "data_ocorencia"`
	ObservationCode            int       `json: "observacao_entrega"`
	Text                       string    `json: "texto"`
	Filler                     string    `json: -`
}

//PROCEDA-3.1
type OccurrenceProceda struct {
	FileName    string `json: "nome_do_arquivo"`
	ContentFile string `json: -`
	HeadFile
	HeadFileTwo
	Carrier
}

//read all content file - OCOREN PROCEDA 3.1
func (proceda *OccurrenceProceda) ReadFile(fileName string) (err error) {
	//Abrir arquivo
	fileOcoren, err := ioutil.ReadFile(fileName)
	checkError(err, "Error: open file")
	proceda.ContentFile = string(fileOcoren)
	originalOcorenSplitLine := strings.Split(proceda.ContentFile, "\n")

	for line := 0; line < len(originalOcorenSplitLine); line++ {
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
			err = proceda.readOccurrences(originalOcorenSplitChar)
			checkError(err, "Error: to read Occurrences")
		case RECORD_CTE:
			err = proceda.dispacherDatas(originalOcorenSplitChar)
			checkError(err, "Error: to read dispacher datas")
		}
	}

	//criar json cabeçalho

	err = proceda.dispacherDatas(originalOcorenSplitLine)
	checkError(err, "Error: to read dispacher Datas")

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
	proceda.HeadFileRecordIdentifier = RECORD_HEARD
	proceda.SenderName = getInformation(originalOcorenSplitChar, SENDER_NAME_INIT, SENDER_NAME_END)
	proceda.RecipientName = getInformation(originalOcorenSplitChar, RECIPIENT_NAME_INIT, RECIPIENT_NAME_END)
	createdAt := getInformation(originalOcorenSplitChar, CREATED_AT_INIT, CREATED_AT_END)
	layout := "2006-01-02T15:04:05.000Z"
	proceda.CreatedAt, err = time.Parse(createdAt, layout)
	checkError(err, "Error: time.Parse(createdAt, layout)")
	return nil
}

const FILE_IDENTIFIER_INIT = 38
const FILE_IDENTIFIER_END = 73

//Cabecalho do arquivo("340")
func (proceda *OccurrenceProceda) readHeadTwo(originalOcorenSplitChar []string) (err error) {
	proceda.HeadFileTwoRecordIdentifier = RECORD_HEARD_TWO
	proceda.FileIdentifier = getInformation(originalOcorenSplitChar, FILE_IDENTIFIER_INIT, FILE_IDENTIFIER_END)
	return nil
}

//"341"
func (proceda *OccurrenceProceda) carrierDatas(originalOcorenSplitChar []string) (err error) {
	/*
		type Carrier struct {
			CarrierRecordIdentifier int    `json: "identificador"`
			RegisteredNumber        string `json: "cnpj_transportadora"`
			Name                    string `json: "nome_transportadora"`
			Filler                  string `json: "-"`
			TransportKnowledges     []TransportKnowledge
		}
	*/
	proceda.CarrierRecordIdentifier = getRecordIdentifier(originalOcorenSplitChar)
	proceda.RegisteredNumber = getRegisteredNumber(originalOcorenSplitChar)
	proceda.Name = getName(originalOcorenSplitChar)
	return nil
}

func getRegisteredNumber(originalOcorenSplitChar []string) (registeredNumber string) {
	registeredNumberInit := 3
	registeredNumberEnd := 13
	for i := registeredNumberInit; i < registeredNumberEnd; i++ {
		registeredNumber = registeredNumber + originalOcorenSplitChar[i]
	}
	return registeredNumber
}
func getName(originalOcorenSplitChar []string) (name string) {
	nameInit := 3
	nameEnd := 13
	for i := nameInit; i < nameEnd; i++ {
		name = name + originalOcorenSplitChar[i]
	}
	return name
}

//"343"
func (proceda *OccurrenceProceda) dispacherDatas(originalOcorenSplitChar []string) (err error) {
	return nil
}

//"342"
func (proceda *OccurrenceProceda) readOccurrences(originalOcorenSplitChar []string) (err error) {
	return nil
}
func checkError(err error, message string) {
	if err != nil {
		fmt.Println(message)
		panic(err)
	}
}

func getRecordIdentifier(originalOcorenSplitChar []string) (recordIdentifier int) {
	record := originalOcorenSplitChar[0] + originalOcorenSplitChar[1] + originalOcorenSplitChar[2]
	fmt.Println("record = ", record)
	recordIdentifier, err := strconv.Atoi(record)
	checkError(err, "err: getRecordIdentifier")
	return recordIdentifier
}

func getInformation(originalOcorenSplitChar []string, init int, end int) (information string) {
	for i := init; i < end; i++ {
		information = information + originalOcorenSplitChar[i]
	}
	return information
}
