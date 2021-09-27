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

//Cabecalho de intercambio ("000") e cabecalho do arquivo("340")
func (proceda *OccurrenceProceda) readHead(originalOcorenSplitChar []string) (err error) {
	proceda.HeadFileRecordIdentifier = RECORD_HEARD
	proceda.SenderName = getSenderName(originalOcorenSplitChar)
	proceda.RecipientName = getRecipientName(originalOcorenSplitChar)
	proceda.CreatedAt = getCreatedAt(originalOcorenSplitChar)

	return nil
}

func getSenderName(originalOcorenSplitChar []string) (senderName string) {
	senderNameInit := 3
	senderNameEnd := 38
	for i := senderNameInit; i < senderNameEnd; i++ {
		senderName = senderName + originalOcorenSplitChar[i]
	}
	return senderName
}
func getRecipientName(originalOcorenSplitChar []string) (recipientName string) {
	recipientNameInit := 38
	recipientNameEnd := 73
	for i := recipientNameInit; i < recipientNameEnd; i++ {
		recipientName = recipientName + originalOcorenSplitChar[i]
	}
	return recipientName
}
func getCreatedAt(originalOcorenSplitChar []string) (createdAt time.Time) {
	return
}

//Cabecalho do arquivo("340")
func (proceda *OccurrenceProceda) readHeadTwo(originalOcorenSplitChar []string) (err error) {
	proceda.HeadFileTwoRecordIdentifier = RECORD_HEARD_TWO
	proceda.FileIdentifier = getFileIdentifier(originalOcorenSplitChar)
	return nil
}

func getFileIdentifier(originalOcorenSplitChar []string) (fileIdentifier string) {
	fileIdentifierInit := 38
	fileIdentifierEnd := 73
	for i := fileIdentifierInit; i < fileIdentifierEnd; i++ {
		fileIdentifier = fileIdentifier + originalOcorenSplitChar[i]
	}
	return fileIdentifier
}

//"341"
func (proceda *OccurrenceProceda) carrierDatas(originalOcorenSplitChar []string) (err error) {
	return nil
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
